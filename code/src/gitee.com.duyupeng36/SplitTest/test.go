package main

import (
	"fmt"
	"gitee.com.duyupeng36/Split"
)

func main() {
	ret := Split.Split("abcdabcdabcd", "b")
	fmt.Println(ret)
	ret = Split.Split("ccccccc", "b")
	fmt.Println(ret)
}
