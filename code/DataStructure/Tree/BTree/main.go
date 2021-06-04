package main

import (
	"fmt"
)

type ElemType rune

type BTree struct {
	root *BNode
}

type BNode struct {
	data   ElemType
	lchild *BNode
	rchild *BNode
}

// CreateBTree 通过括号表示法创建二叉树
func CreateBTree(t *BTree, str string) {
	var stack = make([]*BNode, len([]rune(str)))
	var p *BNode
	var top = -1
	var k int
	for _, ch := range str {
		switch ch {
		case '(':
			top++
			stack[top] = p
			k = 1
		case ')':
			top--
		case ',':
			k = 2
		default:
			p = new(BNode)
			p.data = ElemType(ch)
			p.lchild = nil
			p.rchild = nil
			if t.root == nil {
				(*t).root = p
			} else {
				switch k {
				case 1:
					stack[top].lchild = p
				case 2:
					stack[top].rchild = p
				}
			}
		}
	}

}

// CreateBTreeByPreOrderAndInOrder 通过前序和中序序列创建二叉树
func CreateBTreeByPreOrderAndInOrder(pre, in string) (b *BNode) {
	// pre先序序列 in中序序列 n 为二叉树的结点个数 b构造出二叉树的头节点

	preRune := []rune(pre)
	inRune := []rune(in)

	if len(preRune) == 0 || len(inRune) == 0 {
		return nil
	}

	b = new(BNode)
	b.data = ElemType(preRune[0])
	// 2. 获取根节点在中序遍历数组中的index
	var i int
	for index, value := range inRune {
		if value == preRune[0] {
			i = index
			break
		}
	}
	b.lchild = CreateBTreeByPreOrderAndInOrder(string(preRune[1:i+1]), string(inRune[:i]))
	b.rchild = CreateBTreeByPreOrderAndInOrder(string(preRune[i+1:]), string(inRune[i+1:]))
	return
}

// CreateBTreeByPostOrderAndInOrder 通过后序和中序序列创建二叉树
func CreateBTreeByPostOrderAndInOrder(post, in string) (b *BNode) {
	postRune := []rune(post)
	inRune := []rune(in)

	if len(inRune) < 1 || len(postRune) < 1 {
		return nil
	}
	if len(inRune) == 1 {
		b = new(BNode)
		b.data = ElemType(inRune[0])
		return
	}
	i := 0
	for ; i < len(inRune); i++ {
		if postRune[len(postRune)-1] == inRune[i] {
			break
		}
	}
	b = new(BNode)
	b.data = ElemType(postRune[len(postRune)-1])

	if i == len(postRune)-1 {
		b.lchild = CreateBTreeByPostOrderAndInOrder(string(postRune[:i]), string(inRune[:i]))
	} else if i == 0 {
		b.rchild = CreateBTreeByPostOrderAndInOrder(string(postRune[:len(postRune)-1]), string(inRune[1:]))
	} else {
		b.lchild = CreateBTreeByPostOrderAndInOrder(string(postRune[:i]), string(inRune[:i]))
		b.rchild = CreateBTreeByPostOrderAndInOrder(string(postRune[:len(postRune)-1]), string(inRune[i+1:]))
	}
	return
}

func (t *BTree) DisplayTree(b *BNode) {
	if b != nil {
		fmt.Printf("%c", b.data)
		if b.lchild != nil || b.rchild != nil {
			fmt.Print("(")
			t.DisplayTree(b.lchild)
			if b.rchild != nil {
				fmt.Print(",")
			}
			t.DisplayTree(b.rchild)
			fmt.Print(")")
		}
	}
}

// PreOrder 先序遍历 递归算法
func (t *BTree) PreOrder(b *BNode) {
	if b != nil {
		fmt.Printf("%c ", b.data)
		t.PreOrder(b.lchild)
		t.PreOrder(b.rchild)
	}
}

type SeqStack struct {
	data []*BNode
	top  int
	max  int
}

func InitStack(stack *SeqStack, num int) {
	stack.top = -1
	stack.max = num
	stack.data = make([]*BNode, num)
}
func (s *SeqStack) Push(b *BNode) {
	if s.top > s.max-1 {
		fmt.Println("栈满")
		return
	}
	s.top++
	s.data[s.top] = b
}
func (s *SeqStack) Pop() (b *BNode) {
	if s.top == -1 {
		fmt.Println("栈空")
		return nil
	}
	b = s.data[s.top]
	s.top--
	return
}

// PreOrderNonRecursion 前序遍历的非递归算法
func (t *BTree) PreOrderNonRecursion() {
	var p = t.root
	var stack SeqStack
	InitStack(&stack, 100)
	// 前序遍历的非递归1

	//if p != nil {
	//	stack.Push(p)  // 根节点入栈
	//	// 栈不空，进入循环
	//	for stack.top != -1 {
	//		p = stack.Pop()
	//		// 访问节点
	//		fmt.Printf("%c ", p.data)
	//		// 右子树入栈
	//		if p.rchild != nil {
	//			stack.Push(p.rchild)
	//		}
	//		// 左子树入栈
	//		if p.lchild != nil {
	//			stack.Push(p.lchild)
	//		}
	//	}
	//}

	// 前线遍历的非递归2

	// 栈不为空 或者 根节点不为nil
	for stack.top != -1 || p != nil {
		// 结点不为nil
		for p != nil {
			fmt.Printf("%c ", p.data) // 访问结点
			stack.Push(p)             // 结点入栈
			p = p.lchild              // 访问左子树
		}
		// 此时栈顶结点以访问没有左孩子，或左孩子已经遍历了
		if stack.top != -1 {
			p = stack.Pop() // 出栈
			p = p.rchild    // 处理右孩子
		}
	}
	fmt.Println()

}

// InOrder 中序遍历
func (t *BTree) InOrder(b *BNode) {
	if b != nil {
		t.InOrder(b.lchild)
		fmt.Printf("%c ", b.data)
		t.InOrder(b.rchild)
	}
}

// InOrderNonRecursion 中序遍历的非递归算法
func (t *BTree) InOrderNonRecursion() {
	var p = t.root
	var stack SeqStack
	InitStack(&stack, 100)
	// 栈不空或者根节点不为nil
	for stack.top != -1 || p != nil {
		for p != nil {
			stack.Push(p)
			p = p.lchild
		}
		// 此时栈顶结点没有左孩子或左孩子已经被遍历
		if stack.top != -1 {
			p = stack.Pop()
			fmt.Printf("%c ", p.data)
			p = p.rchild
		}
	}
	fmt.Println()
}

// PostOrder 后序遍历
func (t *BTree) PostOrder(b *BNode) {
	if b != nil {
		t.PostOrder(b.lchild)
		t.PostOrder(b.rchild)
		fmt.Printf("%c ", b.data)
	}
}

// PostOrderNonRecursion 后序遍历的非递归算法
func (t *BTree) PostOrderNonRecursion() {
	var p = t.root
	var stack SeqStack
	var r *BNode
	var flag bool

	InitStack(&stack, 100)

	if p == nil {
		fmt.Println("树为空树")
		return
	}

	for p != nil {
		stack.Push(p)
		p = p.lchild
	}

	r = nil
	flag = true
	for stack.top != -1 && flag {
		p = stack.data[stack.top] // 获取栈顶结点
		// 以访问右子树
		if p.rchild == r {
			fmt.Printf("%c ", p.data)
			p = stack.Pop()
			r = p
		} else {
			p = p.rchild
			flag = false
		}
	}

	for stack.top != -1 {

		for p != nil {
			stack.Push(p)
			p = p.lchild
		}

		r = nil
		flag = true
		for stack.top != -1 && flag {
			p = stack.data[stack.top] // 获取栈顶结点
			// 以访问右子树
			if p.rchild == r {
				fmt.Printf("%c ", p.data)
				p = stack.Pop()
				r = p
			} else {
				p = p.rchild
				flag = false
			}
		}
	}

	fmt.Println()
}

// 层次遍历算法

type SeqQueue struct {
	data             []*BNode
	front, rear, max int
}

func initQueue(queue *SeqQueue, num int) {
	queue.data = make([]*BNode, num)
	queue.rear = -1
	queue.front = -1
	queue.max = num
}

// IsEmpty 判断队列是否为空
func (q *SeqQueue) IsEmpty() bool {
	return q.rear == q.front
}

// IsFull 判断队列是否满队
func (q *SeqQueue) IsFull() bool {
	return (q.rear+1)%q.max == q.front
}

// Put 入队
func (q *SeqQueue) Put(b *BNode) error {
	if q.IsFull() {
		return fmt.Errorf("the queue is full")
	}

	q.rear = (q.rear + 1) % q.max
	q.data[q.rear] = b
	return nil
}

// Get 出队
func (q *SeqQueue) Get() (b *BNode, err error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("the queue is empty")
	}
	q.front = (q.front + 1) % q.max
	b = q.data[q.front]
	return b, nil
}

func (t *BTree) LevelOrder(b *BNode) {
	var p *BNode
	var queue SeqQueue
	initQueue(&queue, 100)
	err := queue.Put(b)
	if err != nil {
		fmt.Println("入队失败")
		return
	}
	for !queue.IsEmpty() {
		p, err = queue.Get()
		if err != nil {
			fmt.Println("出队失败")
			return
		}
		fmt.Printf("%c ", p.data)

		if p.lchild != nil {
			err = queue.Put(p.lchild)
			if err != nil {
				fmt.Println("入队失败")
				return
			}
		}
		if p.rchild != nil {
			err = queue.Put(p.rchild)
			if err != nil {
				fmt.Println("入队失败")
				return
			}
		}
	}
}

func main() {
	//str := "A(B(D(,G)),C(E,F))"
	var t BTree
	//CreateBTree(&t, str)
	root := CreateBTreeByPostOrderAndInOrder("GDBEFCA", "DGBAECF")
	t.root = root
	t.DisplayTree(t.root)
}
