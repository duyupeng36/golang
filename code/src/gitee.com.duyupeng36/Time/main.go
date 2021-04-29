package main

import (
	"fmt"
	"time"
)

// time包
func main() {
	now := time.Now() // 获取当前本地时间

	fmt.Println(now)            // 2021-04-29 11:12:01.4713189 +0800 CST m=+0.001499801
	fmt.Println(now.Location()) // Local 返回时间的时区和地点信息

	name, offset := now.Zone()
	fmt.Println(name, offset)   // CST 28800
	fmt.Println(now.Unix())     // 时间戳 秒 1619666940
	fmt.Println(now.UnixNano()) // 时间戳 纳秒 1619666940164107500
	fmt.Println(now.Date())     // 年月日 2021 April 29
	fmt.Println(now.Clock())    // 时间 12 1 39

	fmt.Println(now.Year())       // 年  2021
	fmt.Println(now.YearDay())    // 一年的第几天 119
	fmt.Println(now.Month())      // 月  April
	fmt.Println(now.Day())        // 日 29
	fmt.Println(now.Hour())       // 时 12
	fmt.Println(now.Minute())     // 分 17
	fmt.Println(now.Second())     // 秒 41
	fmt.Println(now.Nanosecond()) // 秒内偏移量 549089300
	fmt.Println(now.Weekday())    // 星期 Thursday
	fmt.Println(now.ISOWeek())    // IOS周数 2021 17

	fmt.Println(now.Equal(now))         // 等于 true
	fmt.Println(now.Before(time.Now())) // 小于 true
	fmt.Println(now.After(time.Now()))  // 大于 false

	afterHour := now.Add(time.Hour) // 加 一小时
	fmt.Println(afterHour)

	beforeHour := now.Add(-time.Hour) // 减 一小时
	fmt.Println(beforeHour)

	subTime := now.Sub(beforeHour) // 减法
	fmt.Println(subTime)
	//ret, _ := time.Parse("2006/01/02 15:04", "2021/04/29 10:57")   //2021-04-29 10:57:00 +0000 UTC
	//fmt.Println(ret)
	//
	//ret, _ = time.ParseInLocation("2006/01/02 15:04", "2021/04/29 10:57", time.Local)  //2021-04-29 10:57:00 +0800 CST
	//fmt.Println(ret)
	//
	//ret = time.Unix(179072834, 0)  //1975-09-04 22:27:14 +0800 CST
	//fmt.Println(ret)

	timeString := now.Format("2006-01-02 3:04 PM")
	fmt.Println(timeString) // 2021-04-29 12:53 PM
	loc, _ := time.LoadLocation("Asia/Shanghai")
	stringTime, _ := time.ParseInLocation("2006-01-02 3:04 PM", timeString, loc)
	fmt.Println(stringTime)

	//timer := time.Tick(time.Second)
	//for i := range timer {
	//	fmt.Println(i)
	//}

	sub := time.Date(2021, 4, 29, 14, 0, 0, 0, loc).Sub(now)
	fmt.Println(sub)
	now = time.Now().UTC() // UTC只修改Location。
	nextDay, _ := time.ParseInLocation("2006-01-02 15:05:05", "2021-04-30 13:42:50", time.Local)
	fmt.Println(nextDay.Sub(now))
}
