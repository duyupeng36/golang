package main

import "fmt"

// 插入排序

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

// SelectSort 选择排序
func (a *Array) SelectSort() {
	var i, j, k int
	for i = 0; i < a.length-1; i++ {
		k = i
		for j = i + 1; j < a.length; j++ {
			if a.data[j] < a.data[k] {
				k = j
			}
		}
		if k != i {
			a.data[i], a.data[k] = a.data[k], a.data[i]
		}
	}
}

// swift
func (a *Array) swift(low, high int) {
	i := low
	j := 2 * i
	tmp := a.data[i]
	for j <= high {
		// 获取当前结点的最大子结点
		if j < high && a.data[j] < a.data[j+1] {
			j++
		}
		// 比较当前节点与其最大子结点值的大小
		if tmp < a.data[j] {
			// 小就交换
			a.data[i] = a.data[j]
			i = j
			j = 2 * i
		} else {
			// 不小于就跳出循环
			break
		}
	}
	a.data[i] = tmp
}

func (a *Array) HeapSort() {
	// 构造大根堆
	for i := a.length / 2; i >= 0; i-- {
		a.swift(i, a.length-1)
	}

	// 排序
	for i := a.length - 1; i >= 1; i-- {

		tmp := a.data[0]
		a.data[0] = a.data[i]
		a.data[i] = tmp

		a.swift(0, i-1)
	}
}

func main() {
	array := NewArray(9, 6, 1, 3, 4, 2, 8, 5, 10, 7, 0)
	array.HeapSort()
	array.DisplayArray()
}
