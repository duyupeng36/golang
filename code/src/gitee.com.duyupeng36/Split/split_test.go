package Split

import (
	"reflect"
	"testing"
)

//func TestSplit(t *testing.T) {
//	got := Split("a:b:c:d", ":")
//	want := []string{"a","b","c","d"}
//	// 反射比较
//	if !reflect.DeepEqual(got, want) {
//		fmt.Println("测试不通过")
//		t.Errorf("want: %v,but get: %v\n", want, got)
//	} else {
//		fmt.Println("测试通过")
//	}
//
//}

//func TestSplit(t *testing.T) {
//	// 定义一个测试用例类型
//	type test struct {
//		input string
//		sep   string
//		want  []string
//	}
//	// 定义一个存储测试用例的切片
//	tests := []test{
//		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
//		{input: "沙河有沙又有河", sep: "有", want: []string{"沙河", "沙又", "河"}},
//	}
//	// 遍历切片，逐一执行测试用例
//	for _, tc := range tests {
//		got := Split(tc.input, tc.sep)
//		if !reflect.DeepEqual(got, tc.want) {
//			t.Fatalf("excepted:%#v, got:%#v", tc.want, got)
//		}
//	}
//}

func TestSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片
	tests := map[string]test{
		"1": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"2": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"3": {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"4": {input: "沙河有沙又有河", sep: "有", want: []string{"沙河", "沙又", "河"}},
	}
	// 遍历切片，逐一执行测试用例
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

// fib_test.go
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}
