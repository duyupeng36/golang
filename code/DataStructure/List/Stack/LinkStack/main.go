package main

import "fmt"

// 链式栈

type ElemType int

type LinkStack struct {
	root *Node // 根
	top  *Node // 顶层节点
}

type Node struct {
	val  ElemType // 数据元素
	next *Node    // 下一个节点
	prev *Node    // 上一个节点
}

// 初始化一个空栈
func initStack(s *LinkStack) {
	s.root = nil
	s.top = nil
}

// ClearStack 清空栈元素
func (s *LinkStack) ClearStack() {
	for s.top != nil {
		s.top = s.top.prev
		if s.top == nil {
			break
		}
		s.top.next.prev = nil
		s.top.next = nil
	}
	s.root = nil
}

// Push 入栈
func (s *LinkStack) Push(e ElemType) {
	p := new(Node)
	p.val = e
	p.next = nil
	p.prev = nil

	if s.root == nil && s.top == nil {
		s.root = p
		s.top = p
		return
	}

	s.top.next = p
	p.prev = s.top
	s.top = p
}

// Pop 出栈
func (s *LinkStack) Pop() (e ElemType, err error) {
	if s.top == nil {
		return -1, fmt.Errorf("the stack is empty")
	}
	e = s.top.val
	s.top = s.top.prev
	if s.top == nil {
		s.root = nil
		return e, nil
	}
	s.top.next.prev = nil
	s.top.next = nil
	return e, nil
}

// GetTop 获取栈顶元素
func (s LinkStack) GetTop() (e ElemType, err error) {
	if s.top == nil {
		return -1, fmt.Errorf("the stack is empty")
	}

	e = s.top.val

	return e, nil
}

func main() {
	var s LinkStack
	initStack(&s)

	for i := 1; i < 10; i++ {
		s.Push(ElemType(i))
	}
	s.ClearStack()
	fmt.Println(s.GetTop())

	for i := 1; i < 10; i++ {
		fmt.Println(s.Pop())
	}

}
