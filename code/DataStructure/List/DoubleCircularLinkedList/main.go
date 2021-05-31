package main

import (
	"fmt"
)

type ElemType int

type LinkNode struct {
	value      ElemType
	next, prev *LinkNode
}

// 头插法创建一个双向循环链表
func initList(l *LinkNode, a ...ElemType) {
	l.next = l
	l.prev = l

	var s *LinkNode

	for _, v := range a {
		s = new(LinkNode)
		s.value = v
		s.next = l.next
		l.next = s
		s.prev = l
		l.prev = s
	}
}

// DisplayList 输出链表
func (l *LinkNode) DisplayList() {
	h := l.next
	for h != l {
		fmt.Printf("%v<=>", h.value)
		h = h.next
	}
	fmt.Println()
}

// ListEmpty 判断链表是否为空
func (l *LinkNode) ListEmpty() bool {
	return l.next == l
}

// ListLength 求链表长度
func (l *LinkNode) ListLength() (i int) {
	p := l.next

	for p != l {
		i++
		p = p.next
	}
	return
}

// GetElem 获取指定位置上的值
func (l *LinkNode) GetElem(i int, e *ElemType) (err error) {
	if i < 1 {
		return fmt.Errorf("%d out of range", i)
	}
	j := 0
	p := l.next
	for j < i-1 && p != l {
		j++
		p = p.next
	}
	if p == l {
		return fmt.Errorf("%d out of range", i)
	}
	*e = p.value
	return nil
}

// LocateElem 按元素查找
func (l *LinkNode) LocateElem(e ElemType) (i int) {
	p := l.next
	i++
	for p != l && p.value != e {
		i++
		p = p.next
	}
	return
}

// ListInsert 添加元素
func (l *LinkNode) ListInsert(i int, e ElemType) error {
	if i < 1 {
		return fmt.Errorf("%d out of range", i)
	}
	p := l
	j := 0
	for j < i-1 && p.next != l {
		j++
		p = p.next
	}

	if p.next == l && j != i-1 {
		return fmt.Errorf("%d out of range", i)
	}

	s := new(LinkNode)
	s.value = e
	s.next = p.next
	p.next.prev = s
	p.next = s
	s.prev = p
	return nil
}

// ListDelete 删除节点
func (l *LinkNode) ListDelete(i int, e *ElemType) error {
	if i < 1 {
		return fmt.Errorf("%d out of range", i)
	}
	p := l
	j := 0
	for j < i-1 && p.next != l {
		j++
		p = p.next
	}
	if p == l && j != i-1 {
		return fmt.Errorf("%d out of range", i)
	}

	// p.next 就是要删除的节点
	q := p.next
	*e = q.value

	p.next = q.next
	q.next.prev = p
	return nil
}

func main() {
	var h LinkNode
	initList(&h, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	h.DisplayList()
	//fmt.Println(h.ListLength())
	var e ElemType
	err := h.GetElem(1, &e)
	if err != nil {
		return
	}
	fmt.Println(e)
	fmt.Println(h.LocateElem(2))
	err = h.ListInsert(1, 20)
	if err != nil {
		return
	}
	h.DisplayList()
	h.ListDelete(1, &e)
	fmt.Println(e)
	h.DisplayList()
}
