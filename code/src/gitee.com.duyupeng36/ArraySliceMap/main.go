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
	//a1 := [...]int{1, 3, 5, 7, 8}
	//sum := 0
	//for _, v := range a1 {
	//	sum += v
	//}
	//fmt.Printf("sum = %d\n", sum)
	//
	//for i, v1 := range a1 {
	//	for j := i + 1; j < len(a1); j++ {
	//		if v1+a1[j] == 8 {
	//			fmt.Printf("(%d, %d)", i, j) // (0, 3)(1, 2)
	//		}
	//	}
	//}

	//var slice = []int{1,2,3,4}
	//fmt.Println(slice)  // [1 2 3 4]
	//var slice2 = []string{"北京", "上海", "沙河"}
	//fmt.Println(slice2)  // [北京 上海 沙河]

	//var s1 [] int   // 声明一个切片
	//fmt.Println(s1 == nil)  // true
	//// 在golang中nil代表了pointer, channel, Function, interface, map 或者 slice 的零值，类似与c语言中的空指针.
	//var s2 = []int{1,2,3,4}
	//fmt.Println(s2 == nil)  //false

	//var s1 [] int
	//fmt.Println(len(s1), cap(s1))  // 0 0
	//var s2 = []int{1,2,3,4}
	//fmt.Println(len(s2), cap(s2))  // 4 4

	// 有数组得到切片
	//var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s1 := array[0:8]
	//fmt.Printf("s1:%v,type(s1): %T, len(s1): %d, cap(s1): %d\n", s1, s1, len(s1), cap(s1))// s1:[1 2 3 4 5 6 7 8],type(s1): []int, len(s1): 8, cap(s1): 10
	//
	//s2 := array[:5]
	//s3 := array[5:]
	//s4 := array[:]
	//fmt.Println(s2)  // [1 2 3 4 5]
	//fmt.Println(s3)  // [1 2 3 4 5]
	//fmt.Println(s4)  // [1 2 3 4 5 6 7 8 9 10]

	//var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//slice := array[:5]
	//
	//s1 := slice[3: 5]  // s1的容量是底层数组array的长度减去s1的
	//fmt.Printf("s1: %v, cap(s1): %d\n", s1, cap(s1))  // s1: [4 5], cap(s1): 7
	//
	//slice = array[5:]
	////s2 := slice[:6]  // out of range
	//s2 := slice[:5]
	//fmt.Printf("s2: %v, cap(s2): %d", s2, cap(s2))  // s2: [6 7 8 9 10], cap(s2): 5

	//var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s1 := array[:5:6]
	//fmt.Printf("s1: %v, cap(s1): %d\n", s1, cap(s1)) // s1: [1 2 3 4 5], cap(s1): 6
	//s2 := array[4:6:10]
	//fmt.Printf("s2: %v, cap(s2): %d", s2, cap(s2))  // s2: [5 6], cap(s2): 6

	//slice := make([]int, 10, 20)
	//fmt.Printf("slice: %v, len(slice): %d, cap(slice):%d\n", slice, len(slice), cap(slice))  // slice: [0 0 0 0 0 0 0 0 0 0], len(slice): 10, cap(slice):20
	//

	//var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	//s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	//s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	//
	//fmt.Printf("len(s1): %d, cap(s1): %d, s1 is nil: %t\n", len(s1), cap(s1), s1 == nil)  // len(s1): 0, cap(s1): 0, s1 is nil: true
	//fmt.Printf("len(s2): %d, cap(s2): %d, s2 is nil: %t\n", len(s2), cap(s2), s2 == nil)  // len(s2): 0, cap(s2): 0, s2 is nil: false
	//fmt.Printf("len(s3): %d, cap(s3): %d, s3 is nil: %t\n", len(s3), cap(s3), s3 == nil)  // len(s3): 0, cap(s3): 0, s3 is nil: false

	//s1 := make([]int, 3) //[0 0 0]
	//s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	//s2[0] = 100
	//fmt.Println(s1) //[100 0 0]
	//fmt.Println(s2) //[100 0 0]

	//s := []int{1, 3, 5}
	//
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(i, s[i])
	//}
	//
	//for index, value := range s {
	//	fmt.Println(index, value)
	//}

	//var s []int
	//s = append(s, 1)        // 添加一个元素
	//fmt.Printf("s: %v, len(s): %d, cap(s): %d\n", s, len(s), cap(s))
	//s = append(s, 2, 3, 4)  // 添加三个元素
	//fmt.Printf("s: %v, len(s): %d, cap(s): %d\n", s, len(s), cap(s))
	//s2 := []int{5, 6, 7}
	//s = append(s, s2...)  // 添加另一个切片的元素
	//fmt.Printf("s: %v, len(s): %d, cap(s): %d\n", s, len(s), cap(s))

	//var numSlice []int
	//for i := 0; i < 10; i++ {
	//	numSlice = append(numSlice, i)
	//	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	//}

	//a := []int{1, 2, 3, 4, 5}
	//b := a
	//fmt.Println(a) //[1 2 3 4 5]
	//fmt.Println(b) //[1 2 3 4 5]
	//b[0] = 1000
	//fmt.Println(a) //[1000 2 3 4 5]
	//fmt.Println(b) //[1000 2 3 4 5]

	//a := []int{1, 2, 3, 4, 5}
	//c := make([]int, 5, 5)
	//copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	//fmt.Println(a) //[1 2 3 4 5]
	//fmt.Println(c) //[1 2 3 4 5]
	//c[0] = 1000
	//fmt.Println(a) //[1 2 3 4 5]
	//fmt.Println(c) //[1000 2 3 4 5]

	//// 从切片中删除元素
	//a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	//fmt.Println(a)  // [30 31 32 33 34 35 36 37]
	//// 要删除索引为2的元素
	//a = append(a[:2], a[3:]...)
	//fmt.Println(a)  // [30 31 33 34 35 36 37]

	//a := [...]int{30, 31, 32, 33, 34, 35, 36, 37}
	//s := a[:]
	//fmt.Printf("%p ", &s[0])  // 0xc0000c4040
	//fmt.Println(s, len(s), cap(s))  //  [30 31 32 33 34 35 36 37] 8 8
	//
	//s = append(s[:2], s[3:]...)
	//fmt.Printf("%p ", &s[0])  // 0xc0000c4040
	//fmt.Println(s, len(s), cap(s))  // [30 31 33 34 35 36 37] 7 8
	//fmt.Println(a)  // [30 31 33 34 35 36 37 37]

	//var a = make([]int, 5, 10)
	//for i := 0; i < 10; i++ {
	//	a = append(a, i)
	//}
	//fmt.Println(a)  // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

	//a := []int{3,1,7,5,0}
	//sort.Ints(a)
	//fmt.Println(a) // [0 1 3 5 7]

	//a := [...]int{30, 31, 32, 33, 34, 35, 36, 37}
	//
	//s := a[:]
	//
	//s = append(s[:2], s[3:5]...)
	//fmt.Print(s)
	//fmt.Print(a)

	//var m map[string]int
	//m = make(map[string]int, 2)
	//m["age"] = 18
	//m["money"] = 200
	//fmt.Printf("%#v\n", m) // map[string]int{"age":18, "money":200}
	//
	//value, exists := m["hello"]
	//if exists{
	//	fmt.Println(value, exists)  // value默认为对应类型的零值
	//} else {
	//	fmt.Println("该key不在map中")
	//}

	//scoreMap := make(map[string]int)
	//scoreMap["张三"] = 90
	//scoreMap["小明"] = 100
	//scoreMap["娜扎"] = 60
	//delete(scoreMap, "小明")
	//for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}
	//
	//delete(scoreMap, "hello")  // hello不存在，delete不操作

	//var slice1 []map[string]int
	//slice1 = []map[string]int{{"dyp": 19}, {"dyy": 12}}  // 初始化列表初始化
	//fmt.Println(slice1)
	//var slice2 []map[string]int
	//slice2 = make([]map[string]int, 3)  // 初始化切片
	//slice2[0] = make(map[string]int, 3)  // 初始化map
	//slice2[0]["dyp"] = 19
	//slice2[0]["dyy"] = 12
	//fmt.Println(slice2)

	var sliceMap1 map[string][]int
	sliceMap1 = map[string][]int{"北京": {1, 2, 3}, "上海": {4, 5, 6}} // 初始化列表初始化
	fmt.Println(sliceMap1)

	var sliceMap2 map[string][]int
	sliceMap2 = make(map[string][]int, 3)              // 先初始化map
	sliceMap2["北京"] = make([]int, 0, 3)                // 在初始化切片
	sliceMap2["北京"] = append(sliceMap2["北京"], 1, 2, 3) // 向切片添加值
	fmt.Println(sliceMap2)
}
