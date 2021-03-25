package main

import "fmt"

func main() {
	var name string              // 声明变量
	var age int                  // 声明变量
	var isOk bool                // 声明变量
	fmt.Println(name, age, isOk) // 输出变量的默认值

	var a int = 10 // 初始化
	fmt.Println(a)

	var b, c, d int = 11, 12, 13 // 初始化多个同类型变量
	fmt.Println(b, c, d)

	var e = 30   // 自动类型推到
	var f = 3.14 // 自动类型推到
	fmt.Println(e, f)

	g := true // 短变量声明，只能在函数内部使用
	fmt.Println(g)

	const (
		n1 = iota // 0 const出现iota初始化为0
		n2        // 1
		n3        // 2
		n4        // 3
		n5        // 4
	)
	fmt.Println(n1, n2, n3, n4, n5)

	const (
		b1 = iota // 0
		b2        // 1
		_  = iota // 2 该值被跳过
		b3        // 3
	)
	fmt.Println(b1, b2, b3)

	const (
		c1 = iota // 0
		c2 = 100  // 100
		c3        // 100
		c4 = iota // 3
		c5        // 4
	)
	fmt.Println(c1, c2, c3, c4, c5)

	const (
		d1, d2 = iota + 1, iota + 2 // d1 = 1, d2 = 2  iota=0
		d3, d4 = iota + 1, iota + 2 // d3 = 2, d4 = 3  iota=1
		d5, d6                      // d5 = 3, d6 = 4  iota=2
	)
	fmt.Println(d1, d2, d3, d4, d5, d6)

	const (
		_  = iota
		KB = 1 << (10 * iota) // 2^10
		MB = 1 << (10 * iota) // 2^20
		GB = 1 << (10 * iota) // 2^30
		TB = 1 << (10 * iota) // 2^40
		PB = 1 << (10 * iota) // 2^50
	)
	fmt.Println(KB, MB, GB, TB, PB)
}
