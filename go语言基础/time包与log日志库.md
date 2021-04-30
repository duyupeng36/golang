# time时间标准库
time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。

## 一 time包中定义的类型
### 1.1 time.Time 时间类型
```
type Time struct {
    // 内含隐藏或非导出字段
}
```
`Time`代表一个纳秒精度的时间点。表示时间的变量和字段应为`time.Time`
类型，而不是`*time.Time`类型



### 1.2 time.ParseError 时间字符串解析错误类型
```
type ParseError struct {
    Layout     string
    Value      string
    LayoutElem string
    ValueElem  string
    Message    string
}
```
ParseError描述解析时间字符串时出现的错误。

### 1.3 time.Weekday 周类型
```
type Weekday int  // 代表一周的某一天
```
常量
```
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

### 1.4 time.Month 月类型
```
type Month int  // 代表一年的某个月
```
常量
```
const (
    January Month = 1 + iota
    February
    March
    April
    May
    June
    July
    August
    September
    October
    November
    December
)
```

### 1.5 time.Location 地点时区类型
```
type Location struct {
    // 内含隐藏或非导出字段
}
```
Location代表一个（关联到某个时间点的）地点，以及该地点所在的时区
* 本地时间: `var Local *Location = &localLoc`
* UTC通用时间: `var UTC *Location = &utcLoc`


## 二 获取时间
### 2.1 Date构造时间
```
time.Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location)
```
* 返回一个time.Time类型的时间对象
* 该时间对象的时区为loc、当地时间为
  `year-month-day hour:min:sec + nsec nanoseconds`
  的时间点
* 超出范围的修正示例: `October 32`被修正为`November 1`
* `loc`为`nil`会`panic`

### 2.2 Now 获取当前本地时间
```
ret := time.Now()
```
* 返回当前本地时间

### 2.3 Parse 解析时间格式化字符串
```
ret, err := time.Parse(layout, value string)
```
* 解析一个格式化的时间字符串并返回它代表的时间，
* `err`解析时出现错误 
* `layout`解析时的模板，时间点一定要为`Go`的出生时间`2006`年`1`月`2`号`15`点`04`分`05`秒 `Mon Jan`
  没有格式要求
* `value`时间格式要和`layout`的格式一致。

`layout`在time保证的常量有
```
const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // 使用数字表示时区的RFC822
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // 使用数字表示时区的RFC1123
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // 方便的时间戳
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
```

### 2.4  ParseInLocation 解析时间
```
ret, err := time.ParseInLocation(layout, value string, loc *Location)
```
* 类似`Parse`
  * `Parse`将时间解释为`UTC`时间 
  * `ParseInLocation`将`Location`设置为`loc`
### 2.5  Unix 获取本地时间
```
ret := time.Unix(sec int64, nsec int64)
```
* 创建一个本地时间，对应`sec`和`nsec`表示的`Unix`时间
  （从January 1, 1970 UTC至该时间的 *秒数* 和 *纳秒数*）。

### 2.6 获取时间的示例代码
```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now()  // 获取当前本地时间

	fmt.Println(now) // 2021-04-29 11:12:01.4713189 +0800 CST m=+0.001499801

	ret, _ := time.Parse("2006/01/02 15:04", "2021/04/29 10:57")   //2021-04-29 10:57:00 +0000 UTC
	fmt.Println(ret)

	ret, _ = time.ParseInLocation("2006/01/02 15:04", "2021/04/29 10:57", time.Local)  //2021-04-29 10:57:00 +0800 CST
	fmt.Println(ret)

	ret = time.Unix(179072834, 0)  //1975-09-04 22:27:14 +0800 CST
	fmt.Println(ret)
}
```

## 三 时间类型的方法
### 3.1 Location 获取时区和地点
```
func (t Time) Location() *Location
```
`Location`返回`t`的地点和时区信息。
```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now()  // 获取当前本地时间

	fmt.Println(now) // 2021-04-29 11:12:01.4713189 +0800 CST m=+0.001499801
	fmt.Println(now.Location())  // Local 返回时间的时区和地点信息
}
```
### 3.2 Zone 获取时区名和相对UTC的偏移量
```
func (t Time) Zone() (name string, offset int)
```
计算t所在的时区，返回该时区的规范名（如"CET"）和该时区相对于UTC的时间偏移量（单位秒）
```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now()  // 获取当前本地时间

	fmt.Println(now) // 2021-04-29 11:12:01.4713189 +0800 CST m=+0.001499801
	fmt.Println(now.Location())  // Local 返回时间的时区和地点信息

	name,offset := now.Zone()
	fmt.Println(name, offset)  // CST 28800
}
```
### 3.3 IsZero 判断是否为Time的零值
```
func (t Time) IsZero() bool
```
检查`t`是否代表`Time`零值的时间点，`January 1, year 1, 00:00:00 UTC`。

### 3.4 Local 转为本地时间
```
func (t Time) Local() Time
```
返回采用本地和本地时区，但指向同一时间点的`Time`

### 3.5 UTC 转为UTC时间
```
func (t Time) UTC() Time
```
返回采用`UTC`和零时区，但指向同一时间点的`Time`。

### 3.6 In 转为指定时区的时间
```
func (t Time) In(loc *Location) Time
```
返回采用`loc`指定的地点和时区，但指向同一时间点的`Time`。
如果`loc`为`nil`会`panic`。

### 3.7 Unix 时间戳(秒)
```
func (t Time) Unix() int64
```
* 将t时间转为时间戳, 即从时间点`January 1, 1970 UTC`到时间点`t`所经过的时间（单位秒）。

```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now()  // 获取当前本地时间

	fmt.Println(now.Unix())  // 时间戳 秒 1619666940
}

```

### 3.8 UnixNano 时间戳(纳秒)
```
func (t Time) UnixNano() int64
```
* 将t时间转为时间戳, 即从时间点`January 1, 1970 UTC`到时间点`t`所经过的时间（单位 纳秒）。

```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now()  // 获取当前本地时间
	fmt.Println(now.UnixNano())  // 时间戳 纳秒 1619666940164107500
}
```

### 3.9 Date 获取时间点的年月日
```
func (t Time) Date() (year int, month Month, day int)
```
返回时间点t对应的年、月、日
```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now()  // 获取当前本地时间
	fmt.Println(now.Date())  // 年月日 2021 April 29

}
```
### 3.10 Clock 时间
```
func (t Time) Clock() (hour, min, sec int)
```
返回t对应的那一天的时、分、秒。
```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now()  // 获取当前本地时间
	fmt.Println(now.Clock())  // 时间 12 1 39
}
```
### 3.11 Year 年
```
func (t Time) Year() int
```
返回时间点t对应的年份。

### 3.12 Month 月
```
func (t Time) Month() Month
```
返回时间点t对应那一年的第几月。

### 3.13 YearDay 一年的第几天
```
func (t Time) YearDay() int
```
返回时间点t对应的那一年的第几天，平年的返回值范围`[1,365`]，闰年`[1,366]`。
### 3.14 Day 日期
```
func (t Time) Day() int
```
返回时间点t对应那一月的第几日
### 3.15 Weekday 星期
```
func (t Time) Weekday() Weekday
```
回时间点t对应的那一周的周几
### 3.16 ISOWeek 第几周
```
func (t Time) ISOWeek() (year, week int)
```
ISO 9601标准下的年份和星期编号

### 3.17 Hour 小时
```
func (t Time) Hour() int
```
返回t对应的那一天的第几小时，范围`[0, 23]`

### 3.18 Minute 分钟
```
func (t Time) Minute() int
```
返回t对应的那一小时的第几分种，范围`[0, 59]`。

### 3.19 Second 秒
```
func (t Time) Second() int
```
返回t对应的那一分钟的第几秒，范围`[0, 59]`

### 3.20 Nanosecond 秒内纳秒偏移量
```
func (t Time) Nanosecond() int
```
返回t对应的那一秒内的纳秒偏移量，范围`[0, 999999999]`。

## 四 时间类型的比较
### 4.1 Equal 等于
```
func (t Time) Equal(u Time) bool
```
判断两个时间是否相同，会 *考虑时区* 的影响，
因此不同时区标准的时间也可以正确比较。
本方法和用`t==u`不同，这种方法还会比较地点和时区信息。

### 4.2 Before 之前
```
func (t Time) Before(u Time) bool
```
如果`t`代表的时间点在`u`之前，返回真；否则返回假。
### 4.3 After 之后
```
func (t Time) After(u Time) bool
```
如果`t`代表的时间点在`u`之后，返回真；否则返回假。

### 4.4 时间比较示例
```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
    now := time.Now() // 获取当前本地时间
    
    fmt.Println(now.Equal(now))         // 等于 true
    fmt.Println(now.Before(time.Now())) // 小于 true
    fmt.Println(now.After(time.Now()))  // 大于 false
}
```
## 五 时间运算
### 5.1 Duration 时间间隔类型
```
type Duration int64
```
常量
```
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```
### 5.1 Add 加法(加时间间隔)
```
func (t Time) Add(d Duration) Time
```
返回时间点`t+d`

### 5.2 AddDate 加法(加日期) 
```
func (t Time) AddDate(years int, months int, days int) Time
```
`AddDate`返回增加了给出的年份、月份和天数的时间点`Time`

### 5.3 Sub 减法
```
func (t Time) Sub(u Time) Duration
```
返回一个时间段`t-u`。如果结果超出了`Duration`可以表示的 *最大值/最小值*，
将返回 *最大值/最小值* 。要获取时间点`t-d`（d为Duration），
可以使用`t.Add(-d)`。

### 5.4 Time计算示例
```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now()  // 获取当前本地时间
	
	afterHour := now.Add(time.Hour)  // 加 一小时
	fmt.Println(afterHour)

	beforeHour := now.Add(-time.Hour)  // 减 一小时
	fmt.Println(beforeHour)

	subTime := now.Sub(beforeHour)  // 减法
	fmt.Println(subTime)
}
```

## 六 时间字符串格式化

### 6.1 Format 时间格式化为字符串
```
func (t Time) Format(layout string) string
```
`Format`根据`layout`指定的格式返回`t`代表的时间点的格式化文本表示

`layout`使用`Go`的诞生时间`2006`年`1`月`2`号`15`点`04`分`05`秒
来作为格式化的模板字符串。
* 时间点一定要是该时间点
* 格式可以随意
* 格式化为`12`小时方式，需指定`PM`
* 格式化显示毫秒，示例`2006-01-02 15:04:05.000`

**使用ParseInLocation 与 Parse 解析字符串格式的时间**

```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now()  // 获取当前本地时间
	
	timeString := now.Format("2006-01-02 3:04 PM")  // 转为字符串
	fmt.Println(timeString)  // 2021-04-29 12:53 PM
	loc, _ := time.LoadLocation("Asia/Shanghai")  // 获取时区指针
	stringTime, _ := time.ParseInLocation("2006-01-02 3:04 PM", timeString, loc) // 解析字符串时间
	fmt.Println(stringTime)
}
```


## 七 定时器
使用`time.Tick(d Duration)`来设置定时器，
定时器的本质上是一个通道（`channel`）。
```go
package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	timer := time.Tick(time.Second)  // 定时器
	for i := range timer {
		fmt.Println(i)
	}
}
```

# log标准日志库

无论是软件开发的调试阶段还是软件上线之后的运行阶段，
日志一直都是非常重要的一个环节，我们也应该养成在程序中记录日志的好习惯。

`Go`语言内置的`log`包实现了简单的日志服务

## 一 log标准日志库简单使用(标准logger对象)
`log`包定义了`Logger`类型，该类型提供了一些格式化输出的方法。
本包也提供了一个预定义的“标准”`Logger`，可以通过调用函数
* `Print`系列(`Print|Printf|Println`)
* `Fatal`系列(`Fatal|Fatalf|Fatalln`)
  * `Fatal`系列函数会在写入日志信息后调用`os.Exit(1)`
* `Panic`系列(`Panic|Panicf|Panicln`)
  * `Panic`系列函数会在写入日志信息后`panic`。
  
来使用，比自行创建一个logger对象更容易使用，例如下面的示例代码
直接通过`log`包来使用

```go
package main

import (
	"log"
)

func main() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
```
### 1.1 标准logger的配置
默认情况下的`Logger`只会提供日志的时间信息，
但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。
`log`标准库中为我们提供了定制这些设置的方法。
#### flag选项配置
```
func Flags() int
func SetFlags(flag int)
```
* `log`标准库中的`Flags`函数会返回标准`Logger`的输出配置
* `SetFlags`函数用来设置标准`Logger`的输出配置

`log`包中的选项常量
```
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
```
**使用配置：flag选项配置示例**
```
log.SetFlags(log.Llongfile|log.Lmicroseconds|log.Ldate)
```

#### prefix日志前缀配置
`log`标准库中还提供了关于日志信息前缀的两个方法：
```
func Prefix() string
func SetPrefix(prefix string)
```
* `Prefix`函数用来查看标准`Logger`的输出前缀
* `SetPrefix`函数用来设置输出前缀。

**使用配置：前缀配置示例**
```
log.SetPrefix("prefix")
```
* 前缀方便之后对日志信息进行检索和处理

#### 日志输出位置配置
```
func SetOutput(w io.Writer)
```
`SetOutput`函数用来设置标准`Logger`的输出目的地，
默认是标准错误输出。

**配置输出到文件**
```go
package main

import "fmt"
import "log"
import "os"

func main() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("prefix")
	log.Println("这是一条很普通的日志。")
}
```
如果你要使用标准的`Logger`，我们通常会把上面的配置操作写到`init`函数中
```go
package main

import "fmt"
import "log"
import "os"

func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
```

## 二 创建logger
`log`标准库中还提供了一个创建新`Logger`对象的构造函数`New`，
支持我们创建自己的`Logger`示例。`New`函数的签名
```
func New(out io.Writer, prefix string, flag int) *Logger
```
* `out`设置日志信息写入的目的地
* `prefix`前缀
* `flag`定义日志的属性（时间、文件等等）。

**使用示例**

```go
package main

import "log"
import "os"

func main() {
	Logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	Logger.Println("这是自定义的logger记录的日志。")
}
```
