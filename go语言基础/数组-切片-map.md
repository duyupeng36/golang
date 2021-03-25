# 一 数组

数组是同一种数据类型元素的集合。在`Go`语言中，数组从声明时就确定，
使用时可以修改数组成员，但是数组大小不可变化

## 1.1 数组的定义
```
var 数组名 [数组长度] 数据类型
```
* 数组长度: **必须是常量**，并且**是数组类型的一部分**。一旦定义，长度不能变
    * `[5]int`与`[10]int`是两个不同的类型

```go
package main

import "fmt"

func main()  {
  var a [3]int
  var b [4]int
  fmt.Printf("a: %T b: %T\n", a, b) // a: [3]int b: [4]int
}
```

## 1.2 数组的初始化
数组的初始化也有很多方式
### 方式一
初始化数组时可以**使用初始化列表来设置数组元素的值**
```
var 数组名 = [数组长度]类型{初始化列表}
```

**示例**
```go
package main

import "fmt"

func main()  {
	var a [3]int  // 默认使用类型的零值初始化
 	var b = [4]int{1,2,3,4}  // 使用1,2,3,4初始化数组
	var cityArray = [3]string{"北京", "上海", "深圳"}
	fmt.Println(a, b, cityArray)  // [0 0 0] [1 2 3 4] [北京 上海 深圳]
	a = [3]int{5, 6, 7}
	fmt.Println(a)  // [5 6 7]
}
```
### 方式二
**根据初始值的个数自行推断数组的长度**
```
var 数组名 = [...]类型{初始化列表}
```
**示例**
```go
package main

import "fmt"

func main()  {
	var a [3]int  // 默认使用类型的零值初始化
 	var b = [...]int{1,2,3,4}  // 使用1,2,3,4初始化数组
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(a, b, cityArray)  // [0 0 0] [1 2 3 4] [北京 上海 深圳]
	a = [...]int{5, 6, 7}
	fmt.Println(a)  // [5 6 7]
}
```
### 方式三
指定索引值的方式来初始化数组
```
var 数组名 = [数组长度]类型{index: value}
```
**示例**
```go
package main

import "fmt"

func main()  {
	var a [3]int  // 默认使用类型的零值初始化
	a = [3]int{0: 5, 2:7} // 指定索引初始化
	fmt.Println(a)  // [5 0 7]
}
```

## 1.3 数组的遍历
### 索引遍历
```go
package main

import "fmt"

func main()  {
	cityArray := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(cityArray); i++ {
		city := cityArray[i]
		fmt.Println(city)
	}
}
```

### range遍历
```go
package main

import (
	"fmt"
)

func main()  {
	cityArray := [...]string{"北京", "上海", "深圳"}
	for _, v := range cityArray {
		fmt.Println(v)
	}
	
}
```

## 1.4 多维数组

`Go`语言是支持多维数组的，我们这里以二维数组为例（数组中又嵌套数组）
### 1.4.1 二维数组声明
```
var 数组名 [外层长度][内层长度]数据类型
```
**示例**
```go
package main

import (
	"fmt"
)

func main()  {
	var a [3][2] int  // 定义一个三行两列的二维数组. 一维数组中保持了三个只有2个元素一维数组
	fmt.Println(a)  // [[0 0] [0 0] [0 0]]
}
```

### 1.4.2 二维数组的初始化
#### 方式一
使用初始化列表来设置数组元素的值
```
var 数组名 = [外层长度][内层长度]数据类型{数据列表}
```
示例
```go
package main

import (
	"fmt"
)

func main()  {
	var a [3][2] int  // 定义一个三行两列的二维数组. 一维数组中保持了三个只有2个元素一维数组
	fmt.Println(a)  // [[0 0] [0 0] [0 0]]
	a = [3][2]int{{1,2},{3,4},{5,6}}
	fmt.Println(a)  // [[1 2] [3 4] [5 6]]
}
```
#### 方式二
根据初始值的个数自行推断数组的长度
```
var 数组名 = [...][内层长度]数据类型{数据列表}
```
* 仅支持外层自动推算元素个数

**示例**
```go
package main

import (
	"fmt"
)

func main()  {
	var a [3][2] int  // 定义一个三行两列的二维数组. 一维数组中保持了三个只有2个元素一维数组
	fmt.Println(a)  // [[0 0] [0 0] [0 0]]
	a = [...][2]int{{2,3},{4,5},{6,7}}  // 仅支持外层自动推算元素个数，不支持内层自动推算元素个数
	fmt.Println(a)  // [[2 3] [4 5] [6 7]]
}
```
#### 方式三
指定索引值的方式来初始化数组
```
var 数组名 = [外层长度][内层长度]数据类型{index: {内存数据列表}}
```
**示例**
```go
package main

import (
	"fmt"
)

func main()  {
	var a [3][2] int  // 定义一个三行两列的二维数组. 一维数组中保持了三个只有2个元素一维数组
	fmt.Println(a)  // [[0 0] [0 0] [0 0]]
    a = [3][2]int{0: {1,2}, 2: {5,6}}
    fmt.Println(a)  // [[1 2] [0 0] [5 6]]
}
```

#### 1.4.3 二维数组遍历
```go
package main

import (
	"fmt"
)

func main()  {
	var a [3][2] int  // 定义一个三行两列的二维数组. 一维数组中保持了三个只有2个元素一维数组
	fmt.Println(a)  // [[0 0] [0 0] [0 0]]
	a = [3][2]int{{1,2},{3,4},{5,6}}
	for i:=0; i < 3;i++ {
		for j:=0; j < 2; j++ {
			fmt.Println(a[i][j])
		}
	}
	
	for _, v1 := range a {
		for _, v := range v1{
			fmt.Println(v)
		}
	}
}
```

## 1.5 数组是值类型
数组是值类型，赋值和传参会**复制整个数组**。因此**改变副本的值**，
**不会改变本身的值**

```go
package main

import (
	"fmt"
)

func main()  {
	b := [...]int{1, 2, 3}
	b1 := b
	b1[0]=3

	fmt.Println("b:", b)  // b: [1 2 3]
	fmt.Println("b1:", b1)  // b1: [3 2 3]
}
```

**注意**
1. 数组支持 `==`、`!=` 操作符，因为内存总是被初始化过的。
2. `[n]*T`表示指针数组，`*[n]T`表示数组指针 。

****
**数组练习**
```go
package main

import "fmt"

func main()  {
	a1 := [...]int {1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Printf("sum = %d\n", sum)

    for i, v1 := range a1 {
        for j:=i+1; j < len(a1); j++{
            if v1 + a1[j] == 8{
              fmt.Printf("(%d, %d)", i, j) // (0, 3)(1, 2)
            }
        }
    }
}
```

# 二 切片


# 三 map




