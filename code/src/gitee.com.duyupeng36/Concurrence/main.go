package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sing() {
	defer wg.Done()
	for i := 0; ; i++ {
		//runtime.Gosched()  // 让出CPU时间片
		fmt.Printf("---唱歌---%ds\n", i+1)
		time.Sleep(time.Second)
	}
}

func dance() {
	defer wg.Done()
	for i := 0; ; i++ {
		fmt.Printf("---跳舞---%ds\n", i+1)
		time.Sleep(time.Second)
	}
}

func main() {

	//wg.Add(1)
	//go sing()
	//wg.Add(1)
	//go dance()
	//wg.Wait()
	//runtime.GOMAXPROCS(2)
	//
	//for {
	//	go fmt.Print(1)
	//	fmt.Print(0)
	//}
	fmt.Println(runtime.GOMAXPROCS(2))
}
