package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// 二进制
	var a int = 0b10110
	fmt.Printf("十进制:%d -- 二进制: %b 数据类型: %T\n", a, a, a)

	// 八进制  以0开头 或以0o开头
	var b int = 0o77                                    // 077
	fmt.Printf("十进制:%d -- 八进制: %o 数据类型: %T\n", b, b, b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("十进制:%d -- 十六进制: %x 数据类型: %T\n", c, c, c) // ff
	fmt.Printf("十进制:%d -- 十六进制: %X 数据类型: %T\n", c, c, c) // FF

	// 定义一个浮点型数据
	var money float32 = 29382.53
	fmt.Printf("浮点数: %f 数据类型: %T 最大数值为: %e\n", money, money, math.MaxFloat32) // 浮点数: 29382.529297 数据类型: float32 最大数值为: 3.402823e+38
	var gdp float64 = 29382.53
	fmt.Printf("浮点数: %f 数据类型: %T 最大数值为: %e\n", gdp, gdp, math.MaxFloat64) // 浮点数: 29382.530000 数据类型: float64 最大数值为: 1.797693e+308

	// 复数
	var c1 complex64 // 实部和虚部均为32为
	c1 = 1 + 2i
	fmt.Println(c1)   // (1+2i)
	var c2 complex128 // 实部和虚部均为64为
	c2 = 2 + 3i
	fmt.Println(c2) // (2+3i)

	// 布尔类型
	b1 := true  // bool
	var b2 bool // 默认为false
	fmt.Printf("b1数据类型%T\n", b1)
	fmt.Printf("b2的默认值%v\n", b2)

	// 字符串
	s1 := "hello"
	var s2 string = "你好"
	fmt.Printf("%s\n", s1)
	fmt.Printf("%s\n", s2)

	// 转义字符
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	// 多行字符串
	s3 := `静夜思
作者：李白
床前明月光，疑是地上霜。
举头望明月，低头思故乡
`
	fmt.Printf("%s\n", s3)

	// 字符串的常用方法
	fmt.Printf("字符串s1的长度: %d\n", len(s1))
	fmt.Printf("字符串拼接: %s\n", s1+s2)
	fmt.Printf("字符串拼接: %s\n", fmt.Sprintf("%s %s", s1, s2))
	s4 := fmt.Sprintf("%s %s", s1, s2)

	result := strings.Split(s4, " ") // 返会分隔后的切片
	fmt.Printf("%v\n", result)       // [hello 你好]

	fmt.Printf("字符串是否包含\"hello\": %t\n", strings.Contains(s4, "hello")) // 判断字符串是否包含"hello"
	fmt.Printf("是否以\"he\"开头: %t\n", strings.HasPrefix(s4, "he"))        // 是否以"he"开头
	fmt.Printf("是否以\"好\"结尾: %t\n", strings.HasSuffix(s4, "hao"))        // 是否以"好"结尾
	fmt.Printf("子串位置: %d\n", strings.Index(s4, "你好"))                   // 返回子串第一次出现的索引
	fmt.Printf("子串最后一次出现的位置: %d\n", strings.LastIndex(s4, "l"))         // 返回子串最后一次出现的索引
	fmt.Printf("字符串拼接: %s\n", strings.Join(result, "=="))               // 使用==拼接字符串

	s := "hello沙河"
	fmt.Printf("字符串s的长度%d\n", len(s)) // 11, 但是字符串长是7，返回的是一个字节长度
	for i := 0; i < len(s); i++ {     //byte
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()

	// 字符串修改
	s5 := "hello沙河"
	s6 := []rune(s5) // 将s5强制转为rune切片
	s6[5] = '清'
	fmt.Printf("%s\n", string(s6))

	s7 := "hello"
	s8 := []byte(s7) // 将s7转为byte切片
	s8[0] = 'H'
	fmt.Printf("%s\n", string(s8))
}
