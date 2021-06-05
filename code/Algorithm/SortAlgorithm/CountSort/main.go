package main

import (
	"fmt"
)

// 归并

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

func (a Array) DisplayArray() {
	for i := 0; i < a.length; i++ {
		fmt.Printf("%d ", a.data[i])
	}
	fmt.Println()
}

func (a *Array) max() (max int) {
	max = a.data[0]
	for _, v := range a.data {
		if v > max {
			max = v
		}
	}
	return
}

func (a *Array) CountSort() {
	// 获取最大值
	max := a.max()
	var bucket = make([]int, max+1)

	// 记录每个值出现的次数数
	for i := 0; i < a.length; i++ {
		bucket[a.data[i]]++
	}
	// 辅助数组的索引表示原数组保存的值，该索引位置的值为元素出现的次数
	sortedIndex := 0
	for index, v := range bucket {
		for v > 0 {
			a.data[sortedIndex] = index
			sortedIndex++
			v--
		}
	}
}

func main() {
	array := NewArray(9, 6, 1, 3, 4, 2, 8, 5, 10, 7, 0)
	array.CountSort()
	array.DisplayArray()

}
