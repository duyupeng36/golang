package main

import (
	"fmt"
	"time"
)

//
//// producer 生产者
//func producer(send chan <- int)  {
//
//	for {
//		send <- rand.Intn(100)
//	}
//}
//// consumer 消费者
//func consumer(receive<-chan int) {
//	for  {
//		i := <-receive
//		fmt.Printf("消费者获取数据: %d\n", i)
//		time.Sleep(300 * time.Millisecond)
//	}
//}
func main() {
	// 第一种
	fmt.Printf("当前时间: %v\n", time.Now())
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Printf("现下时间: %v\n", time.Now())
	//
	time.Sleep(time.Second)
	// 第三种
	nowTime := <-time.After(time.Second)
	fmt.Printf("现下时间: %v\n", nowTime)
}
