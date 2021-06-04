package main

import (
	"fmt"
)

const MAX = 100

type ElemType int

type SqList struct {
	data   *[MAX]ElemType
	length int
}

// initList 初始化顺序表
func initList(list *SqList, a ...ElemType) {
	list.data = new([MAX]ElemType)
	list.length = 0
	for i, v := range a {
		list.data[i] = v
		list.length++
	}
}

// SeqSearch 顺序查找
func (s *SqList) SeqSearch(elem ElemType) int {
	var i int
	for i < s.length && s.data[i] != elem {
		i++
	}
	if i >= s.length {
		return -1
	}
	return i + 1
}

func (s *SqList) BinSearch(e ElemType) int {
	var low = 0
	var high = s.length - 1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if e == s.data[mid] {
			return mid + 1
		}
		if e < s.data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	var l SqList
	initList(&l, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	i := l.SeqSearch(10)
	fmt.Println(i)
	i = l.BinSearch(10)
	fmt.Println(i)
}
