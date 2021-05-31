package main

import "fmt"

const MAX = 100

type ElemType int

type ArrayStack struct {
	array *[MAX]ElemType // 底层切片
	top   int            // 栈顶位置
}

// 初始化顺序栈
func initStack(stack *ArrayStack) {
	stack.array = new([MAX]ElemType)
	stack.top = -1
}

// ClearStack 清空栈元素
func (s *ArrayStack) ClearStack() {
	s.top = -1
}

// Push 入栈
func (s *ArrayStack) Push(e ElemType) error {
	if s.top == MAX-1 {
		return fmt.Errorf("the stack is full")
	}
	s.top++
	s.array[s.top] = e
	return nil
}

// Pop 出栈
func (s *ArrayStack) Pop() (e ElemType, err error) {
	if s.top == -1 {
		return -1, fmt.Errorf("the stack is empty")
	}

	e = s.array[s.top]
	s.top--
	return e, nil
}

// GetTop 获取栈顶元素
func (s *ArrayStack) GetTop() (e ElemType, err error) {
	if s.top == -1 {
		return -1, fmt.Errorf("the stack is empty")
	}

	e = s.array[s.top]
	return
}

func main() {
	var a ArrayStack
	initStack(&a)

	for i := 1; i < 10; i++ {
		a.Push(ElemType(i))
	}

	for i := 1; i < 10; i++ {
		fmt.Println(a.Pop())
	}

}
