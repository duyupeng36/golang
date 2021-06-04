package main

import "fmt"

// BinarySearchTree 二叉查找树
type BinarySearchTree struct {
	Root *BinarySearchTreeNode // 树根节点
}

// BinarySearchTreeNode 二叉查找树节点
type BinarySearchTreeNode struct {
	Value int64                 // 值
	Times int64                 // 值出现的次数
	Left  *BinarySearchTreeNode // 左子树
	Right *BinarySearchTreeNode // 右字树
}

// NewBinarySearchTree 初始化一个二叉查找树
func NewBinarySearchTree() (t *BinarySearchTree) {
	t = new(BinarySearchTree)
	t.Root = nil
	return
}

// Add 给树添加结点
func (tree *BinarySearchTree) Add(value int64) {
	// 根结点为空 添加根节点并返回
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}
	// 否则 添加子结点
	tree.Root.addNode(value)
}

// addNode 给结点添加子结点
func (node *BinarySearchTreeNode) addNode(value int64) {
	// 值 比当前结点的值 小
	// 添加在左子结点
	if value < node.Value {
		if node.Left == nil {
			// 左子树为空添加
			node.Left = &BinarySearchTreeNode{Value: value}
		} else {
			// 否则 继续递归寻找
			node.Left.addNode(value)
		}
	} else if value > node.Value {
		// 值 比当前结点的值 大
		// 添加在右子结点
		if node.Right == nil {
			// 右子树为空 添加
			node.Right = &BinarySearchTreeNode{Value: value}
		} else {
			// 右子树不为空 递归寻找
			node.Right.addNode(value)
		}
	} else {
		// 相等记录出现次数
		node.Times++
	}
}

// FindMinValue 查找最小值
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.findMinValue()
}

// findMinValue 递归查询左子树
func (node *BinarySearchTreeNode) findMinValue() *BinarySearchTreeNode {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}
	// 一直左子树递归
	return node.Left.findMinValue()
}

// FindMaxValue 查找最大值
func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.findMaxValue()
}

// findMaxValue 递归查找最大值
func (node *BinarySearchTreeNode) findMaxValue() *BinarySearchTreeNode {
	if node.Right == nil {
		return node
	}

	// 递归查找又子树
	return node.Right.findMaxValue()
}

// Find 查找指定元素
func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.find(value)
}

// find 递归查找结点
func (node *BinarySearchTreeNode) find(value int64) *BinarySearchTreeNode {
	// 指定值 等于当前节点值 返回当前节点指针
	if node == nil || value == node.Value {
		return node
	}
	// 指定值 比当前节点值 小
	if value < node.Value {
		//// 递归查询左子树
		//if node.Left == nil {
		//	// 左子树为空，则没有查找到该值 返回nil
		//	return nil
		//}
		// 否则递归查找左子树
		return node.Left.find(value)
	} else {
		//// 指定值 比当前节点值 大
		//// 递归查找右子树
		//if node.Right == nil{
		//	// 当前节点右子树指针为空，表示没有查找到值 则返回nil
		//	return nil
		//}
		// 否则递归查找右子树
		return node.Right.find(value)
	}
}

// FindParent 查找指定元素的父节点
func (tree *BinarySearchTree) FindParent(value int64) *BinarySearchTreeNode {
	// 空树
	if tree.Root == nil {
		return nil
	}
	// 当前值是根节点，没有父节点
	if tree.Root.Value == value {
		return nil
	}
	return tree.Root.findParent(value)
}

// findParent 查找指定值节点的父节点
func (node *BinarySearchTreeNode) findParent(value int64) *BinarySearchTreeNode {
	// 指定值 比当前节点值 小
	if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		// 左子树的根节点的值刚好等于该值，那么父亲就是现在的node，返回
		if node.Left.Value == value {
			return node
		} else {
			return node.Left.findParent(value)
		}
		// 指定值 比当前节点值 大
	} else {
		if node.Right == nil {
			return nil
		}

		if node.Right.Value == value {
			return node
		} else {
			return node.Right.findParent(value)
		}
	}
}

/*
1. 删除的节点是叶子节点，直接删除即可

2. 被删除的节点只有左子树，将被删除节点的左孩子替代被删除的节点即可

3. 被删除的节点只有右子树，将被删除节点的右孩子替代被删除节点即可

4. 被删除节点同时存在左子树和右子树，将其左子树中最大的节点替代被删除节点即可
   （或右子树中最小的节点替代即可）
*/

func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		return
	}
	node := tree.Root.find(value) // 查找值是否存在
	if node == nil {
		return
	}
	parent := tree.Root.findParent(value) // 查找该值的父节点

	// 根节点 且没有子树
	if parent == nil && node.Left == nil && node.Right == nil {
		tree.Root = nil
		return
	}
	// 存在父节点 但是没有子结点 叶子节点
	if parent != nil && node.Left == nil && node.Right == nil {
		// 删除的节点是父节点的左左子结点 直接删除
		if parent.Left.Value == value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	}
	// 存在父节点 存在左子树 但是不存在右子树
	if parent != nil && node.Left != nil && node.Right == nil {
		// 直接将左孩子节点替代删除节点
		if parent.Left.Value == value {
			parent.Left = node.Left
		} else {
			parent.Right = node.Left
		}
		return
	}
	// 存在父节点 左子树为空，右子树不为空
	if parent != nil && node.Left == nil && node.Right != nil {
		// 直接将右孩子替代算法节点接口
		if parent.Left.Value == value {
			parent.Left = node.Right
		} else {
			parent.Right = node.Right
		}
		return
	}
	// 存在父节点 左子树和右子树均不为
	if parent != nil && node.Left != nil && node.Right != nil {
		// 寻找右子树中最小的节点替代被算法的节点
		// 只需要在右子树的左边寻找即可找到
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		// 把最小的节点删掉
		tree.Delete(minNode.Value)
		// 使用最小节点替代
		// 最小值的节点替换被删除节点
		node.Value = minNode.Value
		node.Times = minNode.Times
		return
	}
}

// MidOrder 中序遍历
func (tree *BinarySearchTree) MidOrder() {
	tree.Root.midOrder()
}

func (node *BinarySearchTreeNode) midOrder() {
	if node == nil {
		return
	}

	// 先打印左子树
	node.Left.midOrder()

	// 按照次数打印根节点
	for i := 0; i <= int(node.Times); i++ {
		fmt.Print(node.Value, " ")
	}

	// 打印右子树
	node.Right.midOrder()
}

func main() {

	var t *BinarySearchTree
	t = NewBinarySearchTree()
	values := []int64{9, 2, 3, 7, 1, 4, 10, 8, 6, 5}
	for _, v := range values {
		t.Add(v)
	}

	t.MidOrder()

}
