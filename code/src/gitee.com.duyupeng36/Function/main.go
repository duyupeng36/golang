package main

import (
	"fmt"
	"time"
)
import "math/rand"

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

func plus(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func minus(a ...int) int {
	return a[0] - a[1]
}

func calculator(option func(...int) int, a ...int) (int, func(...int) int) {
	return option(a...), option
}

func function1(a int, b int) int {

	f := func(x int, y int) int {
		return x + y
	} // 匿名函数

	return f(a, b)
}

func function2(a int, b int) int {

	f := func(x int, y int) int {
		return x + y
	}(a, b) // 定义匿名函数，立即执行
	return f
}

// 该函数就是一个闭包函数
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func factorial(n int64) (result int64) {
	if n < 1 {
		result = 1
	} else {
		result = n * factorial(n-1)
	}
	return
}

var count [1000]int64

func fibolach(x int) int64 {
	if x == 1 || x == 2 {
		return 1
	} else if count[x-1] != 0 && count[x-2] != 0 {
		return count[x-1] + count[x-2]
	} else {
		count[x-1] = fibolach(x - 1)
		count[x-2] = fibolach(x - 2)
		return count[x-1] + count[x-2]
	}
}

func f1(f func()) {
	fmt.Println("this is function f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is function f2")
	fmt.Println(x + y)
}

// 将f2函数在f1函数中进行调用，需要对f2函数进行函数定制

func f3(f func(int, int), x, y int) func() {

	return func() {
		f(x, y)
	}
}

// 也可以对函数f2进行如下包装
func f4(f func(int, int)) func(int, int) func() {
	ret := func(x, y int) func() {
		return func() {
			f(x, y)
		}
	}
	return ret
}

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
	//
	//fmt.Println(a)
	//a = 20
	//fmt.Println(a)
	//
	////fmt.Println(b)  // 无法访问变量b
	//b := 30
	//fmt.Println(b) // 变量定义之后才能访问

	//result, option := calculator(plus, 1,2,3,4)
	//fmt.Println(result, option)
	//result, option = calculator(minus, 5, 4)
	//fmt.Println(result, option)
	//
	//fmt.Println(function1(10, 20))
	//
	//fmt.Println(function2(30, 20))
	//{
	//	fmt.Println("hello 匿名函数1")
	//}
	//
	//func () {
	//	fmt.Println("hello 匿名函数2")
	//}()

	//fmt.Println(factorial(13))
	//fibolach(10)
	//fmt.Println(count[0:10])

	rand.Seed(time.Now().UnixNano())
	// 产生随机数
	fmt.Println(rand.Int())    // 产生整型随机数
	fmt.Println(rand.Intn(10)) // 产生10以内的随机数

	function := f3(f2, 10, 20)
	f1(function)

	f := f4(f2)

	ff := f(10, 20)
	f1(ff)
}

//var a int = 10
