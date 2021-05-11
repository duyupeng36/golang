//package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//func reflectType(x interface{}) {
//	v := reflect.TypeOf(x)  // 获取x的数据类型
//	fmt.Printf("type:%v  kind:%v\n", v.Name(), v.Kind())
//}
//
//func reflectValue(x interface{}) {
//	v := reflect.ValueOf(x)
//	k := v.Kind()
//	switch k {
//	case reflect.Int64:
//		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
//		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
//	case reflect.Float32:
//		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
//		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
//	case reflect.Float64:
//		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
//		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
//	}
//}
//
//
//func reflectSetValue(x interface{}) {
//	v := reflect.ValueOf(x)
//	// 反射中使用 Elem()方法获取指针对应的值
//	if v.Elem().Kind() == reflect.Int64 {
//		v.Elem().SetInt(200)
//	}
//}
//
//func main() {
//	//var a float32 = 3.14
//	//reflectType(a) // type:float32  kind:float32
//	//var b int64 = 100
//	//reflectType(b) // type:int64  kind:int64
//	//type student struct {
//	//
//	//}
//	//
//	//var stu student
//	//reflectType(stu)  // type:student  kind:struct
//	//
//	//// 获取值
//	//reflectValue(a) // type is float32, value is 3.140000
//	//reflectValue(b) // type is int64, value is 100
//	//// 将int类型的原始值转换为reflect.Value类型
//	//c := reflect.ValueOf(10)
//	//fmt.Printf("type c :%T\n", c) // type c :reflect.Value
//	//
//	//fmt.Printf("反射设置值前: %d\n", b)
//	//reflectSetValue(&b)
//	//fmt.Printf("反射设置值后: %d\n", b)
//
//
//	// IsNil和IsValue
//
//	// *int类型空指针
//	var a *int
//	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
//	// nil值
//	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
//	// 实例化一个匿名结构体
//	b := struct{}{}
//	// 尝试从结构体中查找"abc"字段
//	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
//	// 尝试从结构体中查找"abc"方法
//	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
//	// map
//	c := map[string]int{}
//	// 尝试从map中查找一个不存在的键
//	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
//
//}
//
package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	printMethod(stu1)
}