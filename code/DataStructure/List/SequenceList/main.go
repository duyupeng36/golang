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

// ListEmpty 判断顺序表是否为空
func (l *SqList) ListEmpty() bool {
	return l.length == 0
}

// ClearList 清空顺序表
func (l *SqList) ClearList() {
	l.length = 0
}

// DisplayList 输出顺序表的内容
func (l *SqList) DisplayList() {
	for i, v := range *l.data {
		fmt.Printf("%v ", v)
		if i == l.length-1 {
			break
		}
	}
	fmt.Println()
}

// GetElem 按位置查找
func (l *SqList) GetElem(i int, e *ElemType) (err error) {
	if i < 1 || i > l.length {
		return fmt.Errorf("%d out of range", i)
	}
	*e = l.data[i]
	return nil
}

// LocateElem 按元素查找
func (l *SqList) LocateElem(e ElemType) (i int) {
	for i < l.length && l.data[i] != e {
		i++
	}
	if i >= l.length {
		return -1
	}
	return i + 1
}

// ListInsert 按位置插入
func (l *SqList) ListInsert(i int, e ElemType) (err error) {
	if i < 1 || i > l.length+1 {
		return fmt.Errorf("%d out of range", i)
	}
	if l.length == MAX {
		return fmt.Errorf("the sequence list is full")
	}
	i--
	for j := l.length; j > i; j-- {
		l.data[j] = l.data[j-1]
	}
	l.data[i] = e
	l.length++

	return nil
}

// ListDelete 删除元素并返回
func (l *SqList) ListDelete(i int, e *ElemType) (err error) {
	if i < 1 || i > l.length {
		return fmt.Errorf("%d out of range", i)
	}
	i--
	*e = l.data[i]

	for j := i; j < l.length-1; j++ {
		l.data[j] = l.data[j+1]
	}
	l.length--
	return nil
}

func main() {
	var l SqList

	initList(&l, 1, 2, 3, 4, 5)

	l.DisplayList()

	err := l.ListInsert(2, 10)
	if err != nil {
		return
	}
	l.DisplayList()
}
