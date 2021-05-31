package main

import "fmt"

// 顺序队列

const MAX = 10

type ElemType int

type SeqQueue struct {
	data        *[MAX]ElemType // 数据域
	rear, front int            // 指针域
}

func initQueue(queue *SeqQueue) {
	queue.data = new([MAX]ElemType)
	queue.rear = -1
	queue.front = -1
}

// IsEmpty 判断队列是否为空
func (q *SeqQueue) IsEmpty() bool {
	return q.rear == q.front
}

// IsFull 判断队列是否满队
func (q *SeqQueue) IsFull() bool {
	return (q.rear+1)%MAX == q.front
}

// Put 入队
func (q *SeqQueue) Put(e ElemType) error {
	if q.IsFull() {
		return fmt.Errorf("the queue is full")
	}

	q.rear = (q.rear + 1) % MAX
	q.data[q.rear] = e
	return nil
}

// Get 出队
func (q *SeqQueue) Get() (e ElemType, err error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("the queue is empty")
	}
	q.front = (q.front + 1) % MAX
	e = q.data[q.front]
	return e, nil
}

func main() {
	var q SeqQueue

	initQueue(&q)

	for i := 1; i < 10; i++ {
		err := q.Put(ElemType(i))
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	for i := 1; i < 10; i++ {
		n, err := q.Get()
		fmt.Println(n, err)
	}

}
