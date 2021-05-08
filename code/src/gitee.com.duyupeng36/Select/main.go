package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	quite := make(chan bool)

	go func() {
		x, y := 1, 1
		for {
			select {
			case ch <- x:
				x, y = y, x+y
			case <-quite:
				runtime.Goexit() // 退出子goroutine
			}
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case num := <-ch:
			fmt.Printf("产生的斐波拉契数 %d\n", num)
			time.Sleep(300 * time.Millisecond)
		}
	}
	quite <- true // 循环结束，告诉子goroutine结束
}
