# context介绍
在 `Go http`包的`Server`中，每一个请求在都有一个对应的 `goroutine` 去处理。
请求处理函数通常会启动额外的 `goroutine` 用来访问后端服务，比如数据库和`RPC`服务。
用来处理一个请求的 `goroutine` 通常需要访问一些与请求特定的数据，
比如终端用户的身份认证信息、验证相关的`token`、请求的截止时间。 
当一个请求被取消或超时时，所有用来处理该请求的 `goroutine` 都应该迅速退出，
然后系统才能释放这些 `goroutine` 占用的资源。

## 如何通知子goroutine退出

### 全局变量的方式
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var exit bool
func f() {
	defer wg.Done()
	i := 0
	for {
		// 判断是否达到退出条件
		if exit {
			break
		}
		i++
		fmt.Printf("执行时间%v\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	wg.addNode(1)
	go f()
	<-time.After(time.Second * 3)  // 等待3秒
	exit = true  // 退出条件置为真
	wg.Wait()
}
```

**缺陷**
1. 使用全局变量在跨包调用时不容易统一
2. 如果`f`中再启动`goroutine`，就不太好控制了。

### 使用channel方式

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

var exit = make(chan bool)
func f() {
	defer wg.Done()
	i := 0
	for {
		select {
		case <-exit:
			runtime.Goexit()  // 退出当前goroutine
		default:  // 如果没有default该子goroutine将阻塞
		}
		i++
		fmt.Printf("执行时间%v\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	wg.addNode(1)
	go f()
	<-time.After(time.Second * 5)  // 等待3秒
	exit <- true  // 写入退出
	wg.Wait()
}
```
**缺陷**
1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的`channel`
### 使用context方式
```go
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()

	go func(ctx context.Context) {
		i := 0
		for {
			select {
			case <-ctx.Done():
				runtime.Goexit()  // 退出当前goroutine
			default:  // 如果没有default该子goroutine将阻塞
			}
			i++
			fmt.Printf("---2执行时间%v\n", i)
			time.Sleep(time.Second)
		}
	}(ctx)

	i := 0
	for {
		select {
		case <-ctx.Done():
			runtime.Goexit()  // 退出当前goroutine
		default:  // 如果没有default该子goroutine将阻塞
		}
		i++
		fmt.Printf("1执行时间%v\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())  // 返回一个ctx对象和取消函数
	wg.addNode(1)
	go f(ctx)
	<-time.After(time.Second * 5)  // 等待3秒
	cancel()  // 通知子goroutine关闭
	wg.Wait()
}
```
**当子`goroutine`又开启另外一个`goroutine`时，只需要将`ctx`传入即可**

`Go1.7`加入了一个新的标准库`context`，它定义了`Context`类型，
专门用来简化 对于处理单个请求的多个 `goroutine` 之间与请求域的数据、
取消信号、截止时间等相关操作，这些操作可能涉及多个 `API` 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。
它们之间的函数调用链必须传递上下文，或者可以使用`WithCancel`、`WithDeadline`、
`WithTimeout`或`WithValue`创建的派生上下文。**当一个上下文被取消时，
它派生的所有上下文也被取消。**

* `WithCancel`: 当需要通知子`goroutine`退出使用
* `WithDeadline`: 取消时间
* `WithTimeout`: 超时时间
* `WithValue`: 传递值

# Context接口
`context.Context`是一个接口，该接口定义了四个需要实现的方法。具体签名如下：
```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```
* `Deadline`方法需要返回当前`Context`被取消的时间，
  也就是完成工作的截止时间（`deadline`）；
* `Done`方法需要返回一个`Channel`，这个`Channel`会在当前工作完成
  或者上下文被取消之后关闭，多次调用`Done`方法会返回同一个`Channel`；
* `Err`方法会返回当前`Context`结束的原因，它只会在`Done`返回的`Channel`
  被关闭时才会返回非空的值；
    * 如果当前`Context`被取消就会返回`Canceled`错误；
    * 如果当前`Context`超时就会返回`DeadlineExceeded`错误；
* `Value`方法会从`Context`中返回键对应的值，对于同一个上下文来说，
  多次调用`Value` 并传入相同的`Key`会返回相同的结果，
  该方法仅用于传递跨`API`和进程间跟请求域的数据；


## Background()和 TODO()
`Go`内置两个函数：`Background()`和`TODO()`，这两个函数分别返回一个实现了
`Context`接口的`background`和`todo`。我们代码中最开始都是以这两个内置的
上下文对象作为最顶层的`partent context`，衍生出更多的子上下文对象。

`Background()`主要用于`main`函数、初始化以及测试代码中，
作为`Context`这个树结构的最顶层的`Context`，也就是 **根`Context`**。

`TODO()`，它目前还不知道具体的使用场景，如果我们不知道该使用什么`Context`的时候，
可以使用这个。

`background`和`todo`本质上都是`emptyCtx`结构体类型，是一个不可取消，
没有设置截止时间，没有携带任何值的`Context`。

# With系列函数
此外，`context`包中还定义了四个With系列函数

## WithCancel
`WithCancel`返回带有新`Done`通道的父节点的副本。
当调用返回的`cancel`函数 或当 关闭父上下文的`Done`通道时，
将关闭返回上下文的`Done`通道，无论先发生什么情况，取消此上下文将释放与其关联的资源。

`WithCancel`的函数签名如下：
```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```
**示例**
```go
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()

	go func(ctx context.Context) {
		i := 0
		for {
			select {
			case <-ctx.Done():
				runtime.Goexit() // 退出当前goroutine
			default: // 如果没有default该子goroutine将阻塞
			}
			i++
			fmt.Printf("---2执行时间%v\n", i)
			time.Sleep(time.Second)
		}
	}(ctx)

	i := 0
	for {
		select {
		case <-ctx.Done():
			runtime.Goexit() // 退出当前goroutine
		default: // 如果没有default该子goroutine将阻塞
		}
		i++
		fmt.Printf("1执行时间%v\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // 返回一个ctx对象和取消函数
	wg.addNode(1)
	go f(ctx)
	<-time.After(time.Second * 5) // 等待3秒
	cancel()                      // 通知子goroutine关闭
	wg.Wait()
}
```
上面的示例代码中，`gen`函数在单独的`goroutine`中生成整数并将它们发送到返回的通道。 
`gen`的调用者在使用生成的整数之后需要取消上下文，以免`gen`启动的内部`goroutine`
发生泄漏。

## WithDeadline
WithDeadline的函数签名如下：
```go
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
```
`deadline`: 一个确定的过期时间；当调用`cancel`或者时间到达`deadline`或者
父上下文的`Done`通道关闭时，取消此上下文将释放与其关联的资源


```go
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	wg.addNode(1)
	go func(ctx context.Context) {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-ctx.Done():
				runtime.Goexit()  // 退出当前goroutine
			default:  // 如果没有default该子goroutine将阻塞
			}
			i++
			fmt.Printf("---2执行时间%v\n", i)
			time.Sleep(time.Second)
		}
	}(ctx)

	i := 0
	for {
		select {
		case <-ctx.Done():
			runtime.Goexit()  // 退出当前goroutine
		default:  // 如果没有default该子goroutine将阻塞
		}
		i++
		fmt.Printf("1执行时间%v\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	d := time.Now().addNode(time.Second * 5)
	ctx, cancel := context.WithDeadline(context.Background(), d)  // 返回一个ctx对象和取消函数
	defer cancel()  // 任何时候调用cancel
	wg.addNode(1)
	go f(ctx)
	//<-time.After(time.Second * 5)  // 等待3秒
	//cancel()  // 通知子goroutine关闭
	wg.Wait()
}
```

## WithTimeout
WithTimeout的函数签名如下：
```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```
`WithTimeout`返回`WithDeadline(parent, time.Now().addNode(timeout))。`

到都达到超时时间将取消此上下文将释放与其相关的资源，
因此代码应该在此上下文中运行的操作完成后立即调用`cancel`，
通常用于 **数据库或者网络连接** 的超时控制
```go
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	wg.addNode(1)
	go func(ctx context.Context) {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-ctx.Done():
				runtime.Goexit()  // 退出当前goroutine
			default:  // 如果没有default该子goroutine将阻塞
			}
			i++
			fmt.Printf("---2执行时间%v\n", i)
			time.Sleep(time.Second)
		}
	}(ctx)

	i := 0
	for {
		select {
		case <-ctx.Done():
			runtime.Goexit()  // 退出当前goroutine
		default:  // 如果没有default该子goroutine将阻塞
		}
		i++
		fmt.Printf("1执行时间%v\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	//d := time.Now().addNode(time.Second * 5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)  // 返回一个ctx对象和取消函数
	defer cancel()

	wg.addNode(1)
	go f(ctx)
	//<-time.After(time.Second * 5)  // 等待3秒
	//cancel()  // 通知子goroutine关闭
	wg.Wait()
}
```

## WithValue
`WithValue`函数能够将请求作用域的数据与 `Context` 对象建立关系。声明如下：
```go
func WithValue(parent Context, key, val interface{}) Context
```
`WithValue`返回父节点的副本，其中与`key`关联的值为`val`。

仅对 **`API`和进程间传递请求域** 的数据使用上下文值，
而不是使用它来传递可选参数给函数

所提供的 **键必须是可比较** 的，并且不应该是`string`类型或任何其他内置类型，
以避免使用上下文在包之间发生冲突。`WithValue`的用户应该为键定义自己的类型。
为了避免在分配给`interface{}`时进行分配，**上下文键通常具有具体类型`struct{}`**。
或者，**导出的上下文关键变量的静态类型应该是指针或接口**。

```go
package main

import (
	"context"
	"fmt"
	"sync"

	"time"
)

// context.WithValue

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg.addNode(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
```

# 使用Context的注意事项

* 推荐以参数的方式显示传递 `Context`
* 以`Context`作为参数的函数方法，应该把`Context`作为第一个参数。
* 给一个函数方法传递`Context`的时候，不要传递`nil`，如果不知道传递什么，
  就使用`context.TODO()`
* `Context`的`Value`相关方法应该传递请求域的必要数据，不应该用于传递可选参数
* `Context`是线程安全的，可以放心的在多个`goroutine`中传递





