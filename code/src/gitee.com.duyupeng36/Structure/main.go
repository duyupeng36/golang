package main

import (
	"fmt"
)

//类型定义
type NewInt int

//类型别名
type MyInt = int

type Person struct { // 定义结构体类型
	name string
	city string
	age  int
}

func search(person Person) Person {
	fmt.Printf("%s-%s-%d\n", person.name, person.city, person.age)
	return person
}

func main() {
	//var a NewInt
	//var b MyInt
	//fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	//fmt.Printf("type of b:%T\n", b) //type of b:int
	//
	//var person Person  // 定义结构体变量
	//// 结构体变量初始化
	//person.name = "dyp"
	//person.age = 18
	//person.city = "上海"
	//
	//fmt.Printf("%#v\n", person)  // main.Person{name:"dyp", city:"上海", age:18}
	//// 结构体变量初始化
	//var person2 = Person{
	//	name: "dyy",
	//	city: "资阳",
	//	age:  12,
	//}
	//fmt.Printf("%#v\n", person2)
	//// 结构体变量的访问
	//fmt.Printf("%s - %s - %d", person2.name, person2.city, person2.age)  // dyy - 资阳 - 12
	//
	//var person3 = Person{"dhs", "资阳", 50}
	//fmt.Printf("%#v\n", person3)
	//
	//person4 := Person{
	//	name: "dyy",
	//	city: "上海",
	//	age: 12,
	//}
	//
	//person5 := person4
	//
	//fmt.Printf("%#v\n", person4)  // main.Person{name:"dyy", city:"上海", age:12}
	//fmt.Printf("%#v\n", person5)  // main.Person{name:"dyy", city:"上海", age:12}
	//person5.age = 25
	//fmt.Printf("%#v\n", person4)  // main.Person{name:"dyy", city:"上海", age:12}
	//fmt.Printf("%#v\n", person5)  // main.Person{name:"dyy", city:"上海", age:25}

	//var people [3] Person  // 定义结构体数组
	//fmt.Printf("%#v\n", people)
	//people = [3]Person{
	//	{
	//		name: "dyp",
	//		city: "上海",
	//		age: 18,
	//	},
	//	{
	//		name: "dyy",
	//		city: "资阳",
	//		age: 12,
	//	},
	//	{
	//		name: "dhs",
	//		city: "资阳",
	//		age: 45,
	//	},
	//}
	//fmt.Printf("%#v\n", people)
	//
	//var people2 [3]Person = [3]Person{
	//	{
	//		name: "dyp",
	//		city: "上海",
	//		age: 18,
	//	},
	//	{
	//		name: "dyy",
	//		city: "资阳",
	//		age: 12,
	//	},
	//	{
	//		name: "dhs",
	//		city: "资阳",
	//		age: 45,
	//	},
	//}
	//fmt.Printf("%#v\n", people2)

	//var person []Person         // 结构体切片声明，没有内存空间
	//fmt.Printf("%#v\n", person) // []main.Person(nil)
	//// 结构体数组初始化
	//person = []Person{
	//	{
	//		name: "dyp",
	//		city: "上海",
	//		age:  18,
	//	},
	//	{
	//		name: "dyy",
	//		city: "资阳",
	//		age:  12,
	//	},
	//	{
	//		name: "dhs",
	//		city: "资阳",
	//		age:  45,
	//	},
	//}
	//
	//fmt.Printf("%#v\n", person)
	//
	//// 定义时初始化
	//var person2 []Person = []Person{
	//	{
	//		name: "dyp",
	//		city: "上海",
	//		age:  18,
	//	},
	//	{
	//		name: "dyy",
	//		city: "资阳",
	//		age:  12,
	//	},
	//	{
	//		name: "dhs",
	//		city: "资阳",
	//		age:  45,
	//	},
	//}
	//fmt.Printf("%#v\n", person2)
	//
	//// 通过make初始化
	//var person3 []Person
	//person3 = make([]Person, 0, 3)
	//person3 = append(person3,Person{name: "dhs", city: "资阳", age:  45})
	//fmt.Printf("%#v\n", person3)
	//
	//
	//for _, v := range person2 {
	//	fmt.Printf("姓名:%s-居住地:%d-年龄:%d\n", v.name, v.city, v.age)
	//}

	//// 先声明后初始化
	//var personMap map[string]Person   // 声明map变量
	//fmt.Printf("%#v\n", personMap)
	//// 变量初始化
	//personMap = map[string]Person{"杜宇鹏": {"dyp", "上海", 19}}
	//fmt.Printf("%#v\n", personMap)
	//// 添加值
	//personMap["杜宇洋"] = Person{"dyy", "资阳", 12}
	//fmt.Printf("%#v\n", personMap)
	//
	//// 定义时初始化
	//var personMap2 map[string]Person = map[string]Person{"杜宇鹏": {"dyp", "上海", 19}}
	//fmt.Printf("%#v\n", personMap2)
	//
	//// make函数初始化
	//var personMap3 map[string]Person
	//personMap3 = make(map[string]Person)
	//// 添加值
	//personMap3["杜宇洋"] = Person{"dyy", "资阳", 12}
	//fmt.Printf("%#v\n", personMap3)
	//
	//for key, value := range personMap {
	//	fmt.Printf("key: %s, value: %#v\n", key, value)
	//}

	var personMap map[string]Person // 声明map变量
	fmt.Printf("%#v\n", personMap)
	// 变量初始化
	personMap = map[string]Person{"杜宇鹏": {"dyp", "上海", 19}}
	fmt.Printf("%#v\n", personMap)
	// 添加值
	personMap["杜宇洋"] = Person{"dyy", "资阳", 12}

	for _, value := range personMap {
		search(value)
	}
}
