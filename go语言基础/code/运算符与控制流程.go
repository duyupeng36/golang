package main

import "fmt"

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

	//var score int32
	//fmt.Println("输入成绩")
	//fmt.Scanf("%d", &score)
	//if score == 100 {
	//	// 90~100属于一个评级
	//	score = 90
	//}
	//switch score / 10 {
	//case 9:
	//	fmt.Println("A")
	//case 8:
	//	fmt.Println("B")
	//case 7:
	//	fmt.Println("C")
	//case 6:
	//	fmt.Println("D")
	//default:
	//	fmt.Println("E")
	//}

	//var i int32 =  0
	//for ; i <= 100; {
	//	if i % 2 != 0{
	//		fmt.Printf("奇数: %d ", i)
	//	}
	//	i++
	//}
	//fmt.Printf("\n")

	//for i:=100; i < 1000; i++ {
	//	hundred := i / 100   // 百位
	//	ten := (i % 100) / 10  // 十位
	//	one := i % 10  // 各位
	//
	//	if hundred * hundred * hundred + ten * ten * ten + one * one * one == i {
	//		fmt.Printf("%d是水仙花数\n", i)
	//}
	//}
	//
	//for i:=7; i<=97; i++ {
	//	if i % 7 == 0 || i / 10 == 7 || i % 10 == 7 {
	//		// 7的倍数    十位是7          个位是7
	//		fmt.Printf("%d ", i)
	//}
	//}

	//// 打印输出九九乘法表
	//for i:=1; i < 10; i++ {
	//	for j:=1; j < i+1; j++ {
	//		fmt.Printf("%d * %d = %d\t", i, j, i * j)
	//	}
	//	fmt.Printf("\n")
	//}
	//fmt.Printf("\n")
	//// 镜像打印九九乘法表
	//for i:=9; i > 0; i-- {
	//	for j:=1; j < i+1; j++ {
	//		fmt.Printf("%d * %d = %d\t", i, j, i * j)
	//	}
	//	fmt.Printf("\n")
	//}

	//for i:=0; i < 10; i++ {
	//	fmt.Printf("%d ", i)
	//	if i == 5 {
	//		break  // 跳出循环
	//	}
	//}
	//fmt.Printf("\n")
	//
	//BREAKDEMO1:
	//	for i := 0; i < 10; i++ {
	//		for j := 0; j < 10; j++ {
	//			if j == 2 {
	//				break BREAKDEMO1  // 跳出标签指定的语句
	//			}
	//			fmt.Printf("%v-%v\n", i, j)
	//		}
	//	}
	//	fmt.Println("...")

	//for i:=0; i < 10; i++ {
	//	if i % 2 != 0 {
	//		continue
	//	}
	//	fmt.Printf("%d ", i)
	//}

	//forloop1:
	//	for i := 0; i < 5; i++ {
	//		//forloop2:
	//		for j := 0; j < 5; j++ {
	//			if i == 2 || j == 2 {
	//				continue forloop1
	//			}
	//			fmt.Printf("%v-%v\n", i, j)
	//		}
	//	}

	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")

}
