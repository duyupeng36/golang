package main

import "fmt"

func main() {
	//var score int32
	//fmt.Println("输入成绩")
	//fmt.Scanf("%d", &score)  // 格式化输入
	//if score > 90 {  // 条件成立执行
	//	fmt.Println("优秀")
	//} else if score > 80 {  // 条件
	//	fmt.Println("良好")
	//} else if score > 70 {
	//	fmt.Println("你有待努力")
	//} else {
	//	fmt.Println("你必须努力了")
	//}
	//
	//if age:=18; age>=18 { // 只能在`if`语句中访问。出了`if`语句范围, 变量`age`被销毁
	//	fmt.Println("澳门赌场上线了")
	//} else {
	//	fmt.Println("你不能进入该场所")
	//}

	// 0~100中的偶数
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("偶数: %d ", i)
		}

	}
	fmt.Printf("\n\n")
	// 0~100中的奇数
	var i int32 = 0
	for i <= 100 {
		if i%2 != 0 {
			fmt.Printf("奇数: %d ", i)
		}
		i++
	}
	fmt.Printf("\n")

	s := "hello world, 你好"
	for i, v := range s {
		fmt.Printf("索引%d: 值%c\n", i, v)
	}

	// 打印输出九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	// 镜像打印九九乘法表
	for i := 9; i > 0; i-- {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Printf("\n")
	}

	// 位运算符, 作用在二进制位上的操作
	// 5: 101
	// 2: 010

	// & 按位与，同位1取1，有0则取0
	fmt.Printf("5 & 2: %b\n", 5&2)
	// | 按位或, 有1取1
	fmt.Printf("5 | 1: %b\n", 5|2)
	// ^ 异或，不同则取1，同则取0
	fmt.Printf("5 ^ 2: %b\n", 5^2)
	// << 左移指定位数: 左移n位就是乘以2的n次方。
	fmt.Printf("5 << 1: %b\n", 5<<1) // 1010
	// >> 右移指定位数: 右移n位就是除以2的n次方。
	fmt.Printf("5 >> 1: %b\n", 5>>1) // 10
	// ~ 按位取反 ~n = -(n+1)
	fmt.Printf("~5: %b\n", ^5) // ~5 = -(5 + 1)
}
