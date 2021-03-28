package main

import "fmt"

func main() {
	//var a int = 10
	//
	//var p *int  // nil
	//fmt.Println(a, p)
	//
	//p = &a
	//fmt.Printf("a:%d a ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	//fmt.Printf("p:%p type:%T\n", p, p) // b:0xc00001a078 type:*int
	//fmt.Printf("&p:%p\n", &p)  // &p:0xc000006028

	//指针取值
	//a := 10
	//b := &a // 取变量a的地址，将指针保存到b中
	//fmt.Printf("type of b:%T\n", b)  // type of b:*int
	//c := *b // 指针取值（根据指针去内存取值）
	//fmt.Printf("type of c:%T\n", c)  // type of c:int
	//fmt.Printf("value of c:%v\n", c) // value of c:10

	//var a[5]int
	//a = [5]int{1,2,3,4,5}
	//
	//var p*[5]int  // 指向数组的指针 -- 数组指针
	//p = &a
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Println(p[i])
	//}

	//var a *int  // 初始化为 *int(nil)
	//*a = 100  // 找不到地址, 出现错误
	//fmt.Println(*a)

	//var a *int = new(int)
	//fmt.Println(a)  // 0xc0000ac068
	//*a = 100  // 向指针执行的内存地址保存数据
	//fmt.Println(*a) // 100

	//a := make([]int, 3, 10)
	//fmt.Println(a)  // [0 0 0]
	//fmt.Printf("%T\n", a)  // []int

	var a [5]*int // 什么一个长度为5的数组，保存int类型的指针
	var b = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		a[i] = &b[i]
	}
	fmt.Println(a) // [0xc00000a480 0xc00000a488 0xc00000a490 0xc00000a498 0xc00000a4a0]

	for i := 0; i < 5; i++ {
		fmt.Println(*a[i])
	}
}
