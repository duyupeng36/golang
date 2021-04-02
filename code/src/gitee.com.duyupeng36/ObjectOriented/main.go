package main

import "fmt"

type MyInt int

func (a MyInt) add(b MyInt) MyInt {
	return a + b
}

type Score struct {
	chinese int
	math    int
	english int
}

type Student struct { // 定义结构体类型
	name   string
	age    int
	gender string
	Score
}

// 构造函数
func InitStudent(name string, age int, gender string, chinese, math, english int) *Student {
	return &Student{
		name:   name,
		age:    age,
		gender: gender,
		Score: Score{
			chinese: chinese,
			math:    math,
			english: english,
		},
	}
}

func (s Student) hello() {
	fmt.Printf("我叫%s，性别%s，今年%d岁\n", s.name, s.gender, s.age)
}

func (s Student) hi() {
	totalScore := s.chinese + s.math + s.english
	avgScore := totalScore / 3
	fmt.Printf("我叫%s，考试总分为%d，平均分为%d\n", s.name, totalScore, avgScore)
}

// 公共属性
type Person struct { // 定义结构体类型
	name   string
	age    int
	gender string
}

func InitPerson(name, gender string, age int) *Person {
	return &Person{
		name:   name,
		age:    age,
		gender: gender,
	}
}
func (p *Person) SayHello() {
	fmt.Printf("我叫%s, 今年%d岁 我是%s生 ", p.name, p.age, p.gender)
}

// 记者
type Reporter struct {
	*Person // 继承公共属性
	hobby   string
}

func initReporter(name, gender, hobby string, age int) *Reporter {
	return &Reporter{
		Person: InitPerson(name, gender, age),
		hobby:  hobby,
	}
}

func (r *Reporter) SayHello() {
	r.Person.SayHello()
	fmt.Printf("我的爱好是%s\n", r.hobby)
}

// 程序员
type Programmer struct {
	*Person
	workYear int
}

func initProgrammer(name, gender string, age, workYear int) *Programmer {
	return &Programmer{
		Person:   InitPerson(name, gender, age),
		workYear: workYear,
	}
}

func (p *Programmer) SayHello() {
	p.Person.SayHello()
	fmt.Printf("我工作%d年了\n", p.workYear)
}
func main0() {
	//type Object struct {
	//	id int
	//	flag bool
	//}
	//
	//type Person struct { // 定义结构体类型
	//	Object
	//	name string
	//	age  int
	//	gender string
	//}
	//
	//type Student struct {
	//	Person  // 匿名字段，继承Person结构体中属性和方法
	//	id int
	//	score int
	//	name string  // 重名字段
	//}
	//
	//type Driver struct {
	//	*Person  // 指针匿名字段，继承Person结构体中属性和方法
	//	id int
	//}
	//
	//var driver Driver
	//driver.Person = new(Person)  // 需要手动初始化指针匿名字段的内存
	//driver.name = "dyp"
	//driver.id = 103
	//driver.Person.name = "dyd"  // error
	//driver.gender = "男"
	//driver.age = 25
	//fmt.Println(driver)  // {0xc000142390 103 dyp}

	//var a MyInt = 10
	//
	//b := a.add(10)
	//fmt.Println(b)  // 20
	//var person = Person{
	//	name:   "dyp",
	//	age:    25,
	//	gender: "男",
	//}
	//person.info()  //  修改前 name:dyp-age:25-gender:男
	//person.edit("dyy", 17, "女")
	//person.info()  // 修改后 name:dyy-age:17-gender:女

	//student := InitStudent("dyp", 19, "男", 98, 99, 87)
	//student.hello()
	//student.hi()

	reporter := initReporter("dyy", "男", "偷拍", 34)
	programmer := initProgrammer("孙权", "男", 23, 3)
	reporter.SayHello()
	programmer.SayHello()
	type methodType func()
	var hello methodType
	hello = reporter.SayHello
	hello()
}
