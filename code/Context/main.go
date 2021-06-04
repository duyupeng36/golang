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
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
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
	//d := time.Now().addNode(time.Second * 5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 返回一个ctx对象和取消函数
	defer cancel()

	wg.Add(1)
	go f(ctx)
	//<-time.After(time.Second * 5)  // 等待3秒
	//cancel()  // 通知子goroutine关闭
	wg.Wait()
}
