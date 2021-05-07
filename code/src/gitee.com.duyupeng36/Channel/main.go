package main

import (
	"fmt"
	"math/rand"
)

//var ch = make(chan int)  // 无缓冲通道
var ch = make(chan int, 10) // 有缓冲冲通道
//var wg sync.WaitGroup

//func printer(s string) {
//	for _, c := range s {
//		fmt.Printf("%c", c)
//		time.Sleep(300 * time.Millisecond)
//	}
//
//}
//
//func person1() {
//	defer wg.Done()
//
//	printer("hello")
//	ch <- 888  // 写入数据
//}
//
//func person2()  {
//	defer wg.Done()
//
//	<- ch  // 读取数据
//	printer(" world!\n")
//}

func producer0() {
	for {
		i := rand.Intn(100)
		ch <- i
		fmt.Printf("生产者生产数据: %d\n", i)
	}
}

func consumer0() {
	for {
		i := <-ch
		fmt.Printf("消费者消费数据: %d\n", i)
	}
}

func main01() {
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch <- i  // 写入数据
	//		fmt.Printf("子goroutine写入数据: %d\n", i)
	//	}
	//	close(ch)  // 关比通道
	//}()
	//
	//for  {
	//	 // 读取数据
	//	if num, ok := <-ch; ok == true{
	//		fmt.Printf("主goroutine读取数据: %d\n", num)
	//	}else {
	//		num := <- ch
	//		fmt.Printf("关闭后读取到的数据: %d\n", num)
	//		break
	//	}
	//}
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}
