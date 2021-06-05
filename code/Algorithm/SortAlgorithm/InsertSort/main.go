package main

import (
	"fmt"
)

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

func (a *Array) InsertSort() {
	for i := 1; i < a.length; i++ {
		if a.data[i] < a.data[i-1] {
			temp := a.data[i]
			j := i - 1
			for j >= 0 && a.data[j] > temp {
				a.data[j+1] = a.data[j]
				j--
			}
			a.data[j+1] = temp
		}
	}
}

func (a *Array) BinInsertSort() {
	var i, j, low, high, mid, tmp int
	for i = 1; i < a.length; i++ {
		if a.data[i] < a.data[i-1] {
			tmp = a.data[i]
			low = 0
			high = i - 1
			for low <= high {
				mid = (low + high) / 2
				if tmp < a.data[mid] {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}

			for j = i - 1; j >= high+1; j-- {
				a.data[j+1] = a.data[j]
			}
			a.data[high+1] = tmp
		}
	}
}

func (a *Array) ShellSort() {
	var i, j, d, tmp int

	d = a.length / 2 // 增量初始值

	for d > 0 {
		// 对所有组采用直接插入排序
		for i = d; i < a.length; i++ {
			tmp = a.data[i] // 对相隔d位置的一组采用直接插入排序
			j = i - d
			for j >= 0 && tmp < a.data[j] {
				a.data[j+d] = a.data[j]
				j = j - d
			}
			a.data[j+d] = tmp
		}
		d = d / 2 // 缩小增量
	}
}

func (a Array) DisplayArray() {
	for i := 0; i < a.length; i++ {
		fmt.Printf("%d ", a.data[i])
	}
	fmt.Println()
}

func main() {

	array := NewArray(9, 6, 1, 3, 4, 2, 8, 5, 10, 7, 0)
	array.ShellSort()
	array.DisplayArray()
}
