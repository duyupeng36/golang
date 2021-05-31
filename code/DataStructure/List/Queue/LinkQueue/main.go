package main

import "fmt"

// 链式队列

type ElemType int

type LinkQueue struct {
	rear, front *Node
}

type Node struct {
	val  ElemType
	next *Node
}

// initQueue 初始化链式队列
func initQueue(queue **LinkQueue) {
	*queue = new(LinkQueue)
	(*queue).rear = nil
	(*queue).front = nil
}

// Put 入队
func (q *LinkQueue) Put(e ElemType) {
	p := new(Node)
	p.val = e
	p.next = nil

	if q.rear == nil && q.front == nil {
		q.rear = p
		q.front = p
		return
	}
	q.rear.next = p
	q.rear = p
}

// Get 出队
func (q *LinkQueue) Get() (e ElemType, err error) {
	if q.rear == nil && q.front == nil {
		err = fmt.Errorf("the queue is empty")
		return -1, err
	}
	p := q.front
	if p == q.rear {
		q.front = nil
		q.rear = nil
		e = p.val
		return e, nil
	}
	q.front = q.front.next
	p.next = nil
	e = p.val
	return e, nil
}

// IsEmpty 判断队列是否为空
func (q *LinkQueue) IsEmpty() bool {
	return q.rear == nil && q.front == nil
}

func main() {
	var q *LinkQueue
	initQueue(&q)
	for i := 1; i < 10; i++ {
		q.Put(ElemType(i))
	}
	for i := 1; i < 10; i++ {
		fmt.Println(q.Get())
	}
	for i := 1; i < 10; i++ {
		fmt.Println(q.Get())
	}
}
