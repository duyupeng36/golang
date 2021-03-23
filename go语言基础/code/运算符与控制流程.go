package main

import (
	"fmt"
)

func main() {
	//a := 10
	//b := 3
	//c := a / b  // c = 3 不超过真实结果的最大整数
	//d := a % b  // 返回的是 a/b 的余数
	//fmt.Println(c, d)  // 3 1
	//
	//a++  // 只能是后自增
	//b--  // 只能是后自减
	//
	////e = a-- // error
	//fmt.Println(a, b)  // 11 2

	//a := 10
	//b := 3.14
	////c := a * b  // 数据类型不相同，不能进行计算
	//c := float64(a) * b
	//fmt.Println(c)

	//var a int32 = 10
	//var b int64 =  3
	////c := a + b  // 数据类型不相同，不能进行计算
	//c := int64(a) + b
	//fmt.Println(c)

	//fmt.Println(64 / 7, 46 % 7)
	// 107653秒是几天几小时几分几秒
	//t := 107653
	//d := t / 60 / 60 / 24 % 365
	//fmt.Printf("%d天", d)
	//h := t / 60 / 60 % 24
	//fmt.Printf("%d时", h)
	//m := t / 60 % 60
	//fmt.Printf("%d分", m)
	//s := t % 60
	//fmt.Printf("%d秒", s)

	//a := 10
	////a += 5  // a = a + 5
	////a -= 10  // a = a - 10
	////a *= 10  // a = a * 10
	////a /= 5  // a = a / 5
	////a %= 3 // a = a % 3
	//
	//a *= 3 + 2  // a = a * (3 + 2)  a = 50
	//a += 3 * 2  // a = a + (3 * 2)  a = 16
	//a += 5 * a  // a = a + (5 * a)  a = 60
	//fmt.Println(a)
	//
	//a := 10
	//b := 10
	//fmt.Println(a == b)  // 判断a是否等于b，等于返回true，不等于返回false
	//a = 10
	//b = 3
	//fmt.Println(a > b)  // 判断a是否大于b，大于返回ture，否则返回false
	//fmt.Println(a < b)  // 判断a是否小于b，小于返回ture，否则返回false
	//
	//fmt.Println(a <= b) // 判断a是否小于或等于b，小于或者等于返回ture，否则返回false
	//fmt.Println(a >= b) // 判断a是否大于或等于b，大于或者等于返回ture，否则返回false
	//
	//fmt.Println(a != b)  // 判断a是否不等于b，不等于返回ture，否则返回false

	//a := false
	//fmt.Println(!a)  // true          // !非，取反
	//fmt.Println(a || true)  // true   // 逻辑或，遇到true则为true
	//fmt.Println(a && true)  // false  // 逻辑与，遇到false则为false

	//a := 10
	//p := &a
	//fmt.Println(p)  // 0xc00000a0b8
	//fmt.Println(*p)  // 10

	//fmt.Println("输入要计算的年份:")
	//var year int
	//fmt.Scan(&year)
	//fmt.Println((year % 400 == 0) || (year % 4 ==0 && year % 100 != 0))

	//var score int32
	//fmt.Println("输入成绩")
	//fmt.Scanf("%d", &score)
	//if score >= 90 {
	//	fmt.Println("优秀")
	//} else if score >= 80 {
	//	fmt.Println("良好")
	//} else if score >= 70 {
	//	fmt.Println("你有待努力")
	//} else {
	//	fmt.Println("你必须努力了")
	//}

	//switch score/10 {
	//case 9:
	//	fmt.Println("优秀")
	//	fallthrough  // 强制执行后面的case
	//case 8:
	//	fmt.Println("良好")
	//	fallthrough
	//case 7:
	//	fmt.Println("你有待努力")
	//	fallthrough  // 强制执行后面的default
	//default:
	//	fmt.Println("你必须努力了")
	//}

	var score int32
	fmt.Println("输入成绩")
	fmt.Scanf("%d", &score)
	if score == 100 {
		// 90~100属于一个评级
		score = 90
	}
	switch score / 10 {
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
