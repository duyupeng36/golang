//package main
//
//import "fmt"
//
//func f1() int {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//
//func f2() (x int) {
//	defer func() {
//		x++  // x = 5 + 1
//	}()
//	return 5  // x = 5  --> x=6
//}
//
//func f3() (y int) {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//func f4() (x int) {
//	defer func(x int) {
//		x++  // x = 5+1=6
//	}(x) // x = 5
//	return 5  // x=5
//}
//func main() {
//	fmt.Println(f1())
//	fmt.Println(f2())
//	fmt.Println(f3())
//	fmt.Println(f4())
//
//	a := 10
//	func(a int){
//		a++
//	}(a)
//	fmt.Println(a)
//}

package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover() // 拦截异常
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B") // 出现异常
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
