package main

import (
	"fmt"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller(0) failed")
		return
	}
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)

}
