package main // 导入主函数包
import "fmt"

// 主函数，go程序运行的入口，有且只有一个主函数
func main() {
	//fmt.Println("你好")
	//var age int  // 声明一个int类型变量
	//age = 10 // 给变量赋值
	//
	//var name string= "杜宇鹏"
	//fmt.Printf("%s: %d", name, age) // 使用变量，输出变量
	//var width float32 = 2.8
	//var height float32 = 3.4
	//var area float32
	//var c float32
	//area = width * height
	//c = (width + height) * 2
	//fmt.Printf("%.2f * %.2f = %.2f\n", width, height, area)
	//fmt.Printf("(%.2f + %.2f)*2 = %.2f\n", width, height, c)

	//pi := 3.1415926
	//r := 3.3
	//area := pi * r * r
	//
	//fmt.Println(area)
	//a, b := 10, 20
	//fmt.Println(a, b)
	//temp := a
	//a = b
	//b = temp
	//fmt.Println(a, b)
	//a,b = b, a
	//fmt.Println(a, b)
	//
	//a = a + b
	//b = a - b  // a=a+b(原来的), a-b=a; 这就可以将a赋值给b
	//a = a - b  // a=a+b(原来的), a - b(新的b)=b(原来的b)；这样就可以将b赋值给a
	//fmt.Println(a, b)

	//a,b,c := 10, 3.1, "dyp"
	//
	//fmt.Println(a, b, c)
	//
	//fmt.Print(a, b, c)
	//
	//fmt.Printf("\n%d,%f,%s\n", a, b, c)

	//var age int
	//var  name string
	//
	//fmt.Println("输入数据: age name")
	//fmt.Scan(&age, &name)
	//fmt.Printf("%s:%d\n", name, age)
	//
	//fmt.Println("输入数据: age name")
	//fmt.Scanf("%d %s", &age, &name)
	//fmt.Println(name, age)

	//var a bool
	//fmt.Println(a)
	//a = true
	//fmt.Println(a)
	//
	//fmt.Printf("%t", a)

	//var a float32
	//var b float64
	//fmt.Println(a, b)
	//a = 4.131922329783
	//b = 4.131922329783
	//fmt.Println(a, b)
	//fmt.Printf("%f,%f", a, b)

	//var c byte
	//fmt.Printf("%T\n", c)  // 查看数据类型 uint8
	//fmt.Println(c)  // 默认值为0
	//c = 'a'
	//fmt.Printf("%c\n", c)  // a
	//fmt.Printf("%d\n", c)  // 97
	//
	//var r rune
	//fmt.Printf("%T\n", r) // 查看数据类型 int32
	//fmt.Println(r) // 默认值为0
	//r = '你'
	//fmt.Printf("%c\n", r)  // 你
	//fmt.Printf("%d\n", r)  // 20320 Unicode编码

	var s1 string
	var s2 string
	fmt.Println(s1) // 字符串默认值'\0'
	fmt.Println(s2) // 字符串默认值'\0'

	s1 = "你好，"
	s2 = "世界！"
	fmt.Println(s1, s2)

	fmt.Println(s1 + s2) // 字符串拼接
}
