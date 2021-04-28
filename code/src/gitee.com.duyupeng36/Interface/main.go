//package main
//
//import "fmt"
//
//type Mover interface {
//	move()
//}
//
//type dog struct {}
//
//// 值接收者实现的接口
//func (d *dog) move() {
//	fmt.Println("狗会动")
//}
//
//func main() {
//	var x Mover
//	/*
//	var wangcai = dog{} // dog类型
//	x = wangcai         // x不可以接收 dog类型
//	x.move()
//	*/
//
//	var fugui = &dog{}  // *dog类型
//	x = fugui           // x可以接收 *dog类型
//	x.move()
//}

//package main
//
//import "fmt"
//
//// Sayer 接口
//type Sayer interface {
//	say()
//}
//
//// Mover 接口
//type Mover interface {
//	move()
//}
//
//type dog struct {
//	name string
//}
//
//// 实现Sayer接口
//func (d dog) say() {
//	fmt.Printf("%s会叫汪汪汪\n", d.name)
//}
//
//// 实现Mover接口
//func (d dog) move() {
//	fmt.Printf("%s会动\n", d.name)
//}
//
//func main() {
//	var x Sayer
//	var y Mover
//
//	var a = dog{name: "旺财"}
//	x = a
//	y = a
//	x.say()
//	y.move()
//}
package main

import "fmt"

// WashingMachine 洗衣机
type WashingMachine interface {
	wash() // 洗衣接口
	dry()  // 甩干接口
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

type Test1 interface { // 子集
	hello()
}

type Test2 interface { // 超集
	Test1 // 继承了一个接口
	read()
}

func main() {
	var t1 Test1 // 子集类型变量
	var t2 Test2 // 超集类型变量
	t1 = t2
	fmt.Println(t1)
}
