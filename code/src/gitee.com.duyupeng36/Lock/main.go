package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup    // 等待组
var rwMutex sync.RWMutex // 创建读写锁
var count = 0            // 全局变量 模拟共享数据

func readGo(in <-chan int, idx int) {
	defer wg.Done()
	for {
		rwMutex.RLock() // 加读锁
		//num := <-in
		fmt.Printf("---第%dth 读goroutine获取数据%d\n", idx, count)
		time.Sleep(time.Second)
		rwMutex.RUnlock() // 解读锁
	}
}

func writeGo(out chan<- int, idx int) {
	defer wg.Done()
	for {
		num := rand.Intn(100) // 生成随机数
		rwMutex.Lock()        // 加写锁
		count = num
		//out <- num
		fmt.Printf("第%dth 写goroutine生成数据%d\n", idx, num)
		time.Sleep(time.Millisecond * 300) // 放大现象
		rwMutex.Unlock()                   // 解写锁
	}

}

func main0() {
	ch := make(chan int)
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go readGo(ch, i+1)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeGo(ch, i+1)
	}
	wg.Wait()
}
