package main

import "fmt"

// 面向对象计算器实现
// 计算器基类
type ObjectOperate struct {
	num1 int
	num2 int
}

// 加法类
type addClass struct {
	ObjectOperate
}

func (add addClass) plus() int {
	return add.num1 + add.num2
}

// 减法类
type subClass struct {
	ObjectOperate
}

func (sub subClass) minus() int {
	return sub.num2 - sub.num1
}

// Sayer 接口
type Sayer interface {
	say()
}
type dog struct{}

type cat struct{}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

func main() {
	var add addClass
	add.num2 = 10
	add.num1 = 20

	sum := add.plus()
	fmt.Println(sum)

	var sub subClass
	sub.num1 = 10
	sub.num2 = 20
	minus := sub.minus()
	fmt.Println(minus)

	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}
