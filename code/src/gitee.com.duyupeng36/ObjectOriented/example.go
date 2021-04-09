package main

import "fmt"

// ObjectOperate 面向对象计算器实现
// 计算器基类
type ObjectOperate struct {
	num1 int
	num2 int
}

// 加法类
type addClass struct {
	ObjectOperate
}

func (add *addClass) Operate() int {
	return add.num1 + add.num2
}

// 减法类
type subClass struct {
	ObjectOperate
}

func (sub *subClass) Operate() int {
	return sub.num2 - sub.num1
}

// Operation 接口
type Operation interface {
	Operate() int
}

// OptFactory 多态实现
func OptFactory(o Operation) (value int) {
	value = o.Operate()
	return
}

// Factory 工厂模式，空结构体
type Factory struct {
}

func (f *Factory) Calc(number1, number2 int, op string) (value int) {
	// 接口类型变量
	var opt Operation
	switch op {
	case "+":
		var add addClass = addClass{ObjectOperate{number1, number2}} // 创建加法对象
		opt = &add                                                   // 绑定对象
	case "-":
		var sub subClass = subClass{ObjectOperate{number1, number2}} // 创建减法对象
		opt = &sub                                                   // 绑定的对象
	}
	//value = opt.Operate()  // 调用接口
	value = OptFactory(opt)
	return
}

func main() {
	//var add addClass
	//add.num2 = 10
	//add.num1 = 20
	//
	//sum := add.Operate()
	//fmt.Println(sum)
	//
	//var sub subClass
	//sub.num1 = 10
	//sub.num2 = 20
	//minus := sub.Operate()
	//fmt.Println(minus)
	var factory Factory
	value := factory.Calc(10, 20, "+")
	fmt.Println(value)
}
