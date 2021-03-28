package main

import "fmt"

// 定义一个函数，返回两个数的和
func add(x, y int) int {
	return x + y
}

func intSum(a ...int) int {
	sum := 0
	// 传递过来的参数被打包到了一个切片a中
	for _, v := range a {
		sum += v
	}
	return sum
}

func intSum2(sum int, a ...int) int {
	for _, v := range a {
		sum += v
	}
	return sum
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func initSlice(slice *[]int) {
	*slice = make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		*slice = append(*slice, i)
	}
}

func test1() {
	test2()
	fmt.Println("test1")
}

func test2() {
	fmt.Println("test2")
}

func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

type funcType func(int, int) (int, int)

func main() {
	//z := add(10, 20) // 调用函数
	//fmt.Println(z)
	//
	//z = intSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//fmt.Println(z)
	//
	//sum := 10
	//z = intSum2(sum, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//fmt.Println(z)
	//
	//a := 10
	//b := 20
	//fmt.Println("交换之前:", a, b)
	//swap(&a, &b)
	//fmt.Println("交换之后", a, b)
	//
	//var slice []int
	//initSlice(&slice)
	//fmt.Println(slice)

	//test1()

	//x, y := calc(10, 20)
	//fmt.Println(x, y)
	//var ca funcType
	//ca = calc
	//x, y = 30, -10
	//x, y = ca(x, y)
	//fmt.Println(x, y)

	fmt.Println(a)
	a = 20
	fmt.Println(a)

	//fmt.Println(b)  // 无法访问变量b
	b := 30
	fmt.Println(b) // 变量定义之后才能访问

}

var a int = 10
