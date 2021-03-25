package main

import "fmt"

func main() {
	//var a [3]int  // 默认使用类型的零值初始化
	//var b = [4]int{1,2,3,4}  // 使用1,2,3,4初始化数组
	//var cityArray = [...]string{"北京", "上海", "深圳"}
	//fmt.Println(a, b, cityArray)  // [0 0 0] [1 2 3 4] [北京 上海 深圳]
	//a = [3]int{0: 5, 2:7}
	//fmt.Println(a)  // [5 0 7]
	//cityArray := [...]string{"北京", "上海", "深圳"}
	//for i := 0; i < len(cityArray); i++ {
	//	city := cityArray[i]
	//	fmt.Println(city)
	//}
	//
	//for _, v := range cityArray {
	//	fmt.Println(v)
	//}
	//
	//var a [3][2] int  // 定义一个三行两列的二维数组. 一维数组中保持了三个只有2个元素一维数组
	//fmt.Println(a)  // [[0 0] [0 0] [0 0]]
	//a = [3][2]int{{1,2},{3,4},{5,6}}
	////fmt.Println(a)  // [[1 2] [3 4] [5 6]]
	////a = [...][2]int{{2,3},{4,5},{6,7}}  // 仅支持外层自动推算元素个数，不支持内层自动推算元素个数
	////fmt.Println(a)  // [[2 3] [4 5] [6 7]]
	////a = [3][2]int{0: {1,2}, 2: {5,6}}
	////fmt.Println(a)  // [[1 2] [0 0] [5 6]]
	//
	//for i:=0; i < 3;i++ {
	//	for j:=0; j < 2; j++ {
	//		fmt.Println(a[i][j])
	//	}
	//}
	//
	//for _, v1 := range a {
	//	for _, v := range v1{
	//		fmt.Println(v)
	//	}
	//}
	//
	//b := [...]int{1, 2, 3}
	//b1 := b
	//b1[0]=3
	//
	//fmt.Println("b:", b)  // b: [1 2 3]
	//fmt.Println("b1:", b1)  // b1: [3 2 3]
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Printf("sum = %d\n", sum)

	for i, v1 := range a1 {
		for j := i + 1; j < len(a1); j++ {
			if v1+a1[j] == 8 {
				fmt.Printf("(%d, %d)", i, j) // (0, 3)(1, 2)
			}
		}
	}
}
