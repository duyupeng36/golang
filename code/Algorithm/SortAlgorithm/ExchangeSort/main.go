package main

import (
	"fmt"
)

// 交换排序

type Array struct {
	data   []int
	length int
}

const MAX = 100

func NewArray(a ...int) (array *Array) {
	array = &Array{
		data:   make([]int, MAX),
		length: 0,
	}

	for i, v := range a {
		array.data[i] = v
		array.length++
	}
	return
}

// bubbleSort 冒泡排序
func (a *Array) bubbleSort() {
	var i, j int
	var isExcjange bool
	for i = 0; i < a.length-1; i++ {
		// 开始循环没有发生交换
		isExcjange = false

		//// 前面一个元素比后面一个元素关键字大时，交换两个元素
		//for j = 0; j < a.length - i - 1; j++ {
		//	if a.data[j] > a.data[j+1] {
		//		temp := a.data[j]
		//		a.data[j] = a.data[j+1]
		//		a.data[j+1]=temp
		//		// 发生交换
		//		isExcjange=true
		//	}
		//}
		//

		// 后一个元素关键字比前一个元素的关键字小时，交换两个元素
		for j = a.length - 1; j > i; j-- {
			if a.data[j] < a.data[j-1] {

				//temp := a.data[j]
				//a.data[j] = a.data[j - 1]
				//a.data[j-1] = temp

				// 交换两个元素
				a.data[j], a.data[j-1] = a.data[j-1], a.data[j]
				isExcjange = true
			}
		}

		// 发生了交换，循环不退出，如果没有发送交换循环退出
		if !isExcjange {
			break
		}
	}
}

// partition 分区
func (a *Array) partition(s, t int) int {
	// s 排序区间的起始位置
	i := s
	j := t
	temp := a.data[i]
	for i < j {
		for j > i && a.data[j] >= temp {
			j--
		}
		a.data[i] = a.data[j]
		for i < j && a.data[i] <= temp {
			i++
		}
		a.data[j] = a.data[i]
	}
	a.data[i] = temp
	return i
}

// QuickSort 快速排序
func (a *Array) QuickSort(s, t int) {

	piovt := (s + t) / 2 // 基准位置选择序列最中间的位置

	if s < t {
		if piovt != s {
			temp := a.data[piovt]
			a.data[piovt] = a.data[s]
			a.data[s] = temp
		}
		i := a.partition(s, t)
		a.QuickSort(s, i-1)
		a.QuickSort(i+1, t)
	}
}

func (a Array) DisplayArray() {
	for i := 0; i < a.length; i++ {
		fmt.Printf("%d ", a.data[i])
	}
	fmt.Println()
}

func main() {
	array := NewArray(9, 8, 7, 6, 5, 4, 3, 2, 1, 0)
	//array.bubbleSort()
	array.QuickSort(0, array.length-1)
	array.DisplayArray()
}
