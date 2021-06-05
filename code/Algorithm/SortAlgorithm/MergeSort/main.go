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

// Merge 一次归并
func (a *Array) merge(low, mid, high int) {
	// 申请辅助数组
	auxiliaryArray := make([]int, high-low+1)
	var i = low
	var j = mid + 1
	var k = 0
	// 两个子序列均为被扫描完
	for i <= mid && j <= high {
		// 前一个子序列中的值，比后一个子序列中的值小或者等于时
		if a.data[i] <= a.data[j] {
			auxiliaryArray[k] = a.data[i] // 将第一个子序列中的值放入辅助数组
			i++
			k++
		} else {
			auxiliaryArray[k] = a.data[j] // 否则将第二个子序列中的值放入辅助数组
			j++
			k++
		}
	}
	// 第二个子序列已经放完，将第一个子序列余下的值放入辅助数组
	for i <= mid {
		auxiliaryArray[k] = a.data[i]
		i++
		k++
	}
	// 第一个子序列已经放完，将第二个子序列余下的值放入辅助数组
	for j <= high {
		auxiliaryArray[k] = a.data[j]
		j++
		k++
	}

	// 最后将辅助数组中的值复制到原始数组中
	k = 0
	for i = low; i <= high; i++ {
		a.data[i] = auxiliaryArray[k]
		k++
	}
}

// MergeSort 自顶向下的二路归并
func (a *Array) MergeSort(low, high int) {
	var mid int
	if low < high {
		mid = (low + high) / 2
		a.MergeSort(low, mid)
		a.MergeSort(mid+1, high)
		a.merge(low, mid, high)
	}
}

func main() {
	array := NewArray(9, 6, 1, 3, 4, 2, 8, 5, 10, 7, 0)
	array.MergeSort(0, array.length-1)
	array.DisplayArray()
}
