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

	//var a [5]*int // 什么一个长度为5的数组，保存int类型的指针
	//var b = [5]int{1, 2, 3, 4, 5}
	//
	//for i := 0; i < 5; i++ {
	//	a[i] = &b[i]
	//}
	//fmt.Println(a) // [0xc00000a480 0xc00000a488 0xc00000a490 0xc00000a498 0xc00000a4a0]
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Println(*a[i])
	//}

	//var slice []int = []int{1,2,3,4,5}
	//var p *[]int = &slice
	//fmt.Printf("%p\n", p)  // 0xc000118060
	//fmt.Printf("%p\n", *p)  // 0xc000156060
	//fmt.Printf("%p\n", slice) // 0xc000156060
	//fmt.Printf("%p\n", &slice[0]) // 0xc000156060
	//fmt.Println((*p)[0])  // 通过切片指针获取切片底层数组保存的值
	//*p = append(*p, 6, 7, 8, 9, 10,11,12,13,14,15,16,17,18)
	//// *p == slice  修改切片对应的地址
	//fmt.Println(*p)
	//fmt.Println(slice)
	//fmt.Printf("%p\n", slice)
	//var a[5]int  // 声明一个数组
	//a = [5]int{1,2,3,4,5}  // 初始化
	//
	//var p*[5]int  // 声明一个指向数组的指针 -- 数组指针
	//p = &a  // 初始化
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Println((*p)[i])
	//	fmt.Println(p[i])  // 与数组本身操作没有任何区别
	//}
	//
	//fmt.Printf("%p\n", &a)  // 0xc0000d6060
	//fmt.Printf("%p\n", &a[0]) // 0xc0000d6060
	//fmt.Printf("%T\n", &a)  // *[5]int
	//fmt.Printf("%T\n", &a[0])  // *int
	//// &a与&a[0]虽然是同一个个地址，但是类型是不一样的。

	//a := 10
	//b := 20
	//var p []*int = []*int{&a, &b}
	//fmt.Println(p)  // [0xc00000e0d8 0xc00000e100]
	//
	//type Person struct { // 定义结构体类型
	//	name string
	//	city string
	//	age  int
	//}
	//var person Person = Person{"你好", "上海",12}
	//
	//fmt.Printf("%#v\n", person)
	//var p * Person
	//p = &person
	//fmt.Printf("%p\n", p)  // 0xc0000743c0
	//fmt.Printf("%p\n", &person.name)  // 0xc0000743c0
	//fmt.Printf("%s\n", p.name)  // 通过结构体指针操作结构体成员
	//fmt.Printf("%s\n", p.city)
	//fmt.Printf("%d\n", p.age)

	a := 10

	p1 := &a
	p2 := &p1
	fmt.Printf("%p\n", &p1) // 0xc0000d8018
	fmt.Printf("%p\n", p2)  // 0xc0000d8018
	fmt.Println(**p2)
}
