package main

import "fmt"

// 二叉查找树
type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

// 二叉查找树节点
type BinarySearchTreeNode struct {
	Value int64                 // 值
	Times int64                 // 值出现的次数
	Left  *BinarySearchTreeNode // 左子树
	Right *BinarySearchTreeNode // 右子树
}

// 一个节点代表一个元素，节点的 Value 值是用来进行二叉查找的关键
// 当 Value 值重复时，我们将值出现的次数 Times + 1

func (tree *BinarySearchTree) Add(value int64) {
	// 如果没有树根，证明是棵空树，添加树根后返回
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}
	tree.Root.Add(value)
}

// 将值添加进去
func (node *BinarySearchTreeNode) Add(value int64) {
	// 如果插入的值比节点的值小，那么要插入到该节点的左子树
	if value < node.Value {
		// 如果左子树为空，直接添加
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{Value: value}
		} else {
			// 否则递归
			node.Left.Add(value)
		}
	} else if value > node.Value {
		// 如果插入的值比节点的值大，那么要插入到该节点的右子树
		// 如果右子树为空，直接添加
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: value}
		} else {
			// 否则递归
			node.Right.Add(value)
		}
	} else {
		// 值相同，不需要添加，值出现的次数 + 1
		node.Times = node.Times + 1
	}
}

// 查找最大值或最小值的元素

// 找出最小值的节点
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	// 左子树为空，表面已经是最左的节点，该值就是最小值
	if node.Left == nil {
		return node
	}
	// 一直左子树递归
	return node.Left.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	// 右子树为空，表面已经是最右的节点，该值就是最大值
	if node.Right == nil {
		return node
	}
	// 一直右子树递归
	return node.Right.FindMaxValue()
}

// 二分查找
// 查找节点
func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.Find(value)
}

func (node *BinarySearchTreeNode) Find(value int64) *BinarySearchTreeNode {
	if value == node.Value {
		// 找到该节点了
		return node
	} else if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			// 左子树为空，表示找不到该值了，返回 nil
			return nil
		}
		return node.Left.Find(value)
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		if node.Right == nil {
			// 右子树为空，表示找不到该值了，返回 nil
			return nil
		}
		return node.Right.Find(value)
	}
}

// 查找指定元素的父亲
func (tree *BinarySearchTree) FindParent(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	// 如果根节点等于该值，根节点没有父节点，返回 nil
	if tree.Root.Value == value {
		return nil
	}
	return tree.Root.FindParent(value)
}

func (node *BinarySearchTreeNode) FindParent(value int64) *BinarySearchTreeNode {
	// 外层无值相等判定，在内层判断完毕
	if value < node.Value {
		leftTree := node.Left
		if leftTree == nil {
			// 左子树为空，找不到该值了，返回nil
			return nil
		}

		// 左子树根节点的值等于该值，父节点为现在的节点，返回该值
		if leftTree.Value == value {
			return node
		} else {
			return leftTree.FindParent(value)
		}
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始
		rightTree := node.Right
		if rightTree == nil {
			// 右子树为空，找不到该值了，返回 nil
			return nil
		}

		// 右子树根节点的值等于该值，父节点为现在的节点，返回该值
		if rightTree.Value == value {
			return node
		} else {
			return rightTree.FindParent(value)
		}
	}
}

// 删除元素
// 删除元素4种情况
// 1.删除根节点且根节点没有儿子,直接删除
// 2.删除节点有父节点，但没哟子树，直接删除
// 3.删除节点右两个子树，右子树比左子树大，用右子树中最小元素来替换删除的节点
// 4.删除的节点只有一个子树，直接替换被删除的节点即可

// 删除指定元素
func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		// 如果是空树，直接返回
		return
	}

	// 查找该值是否存在
	node := tree.Root.Find(value)
	if node == nil {
		// 该值不存在
		return
	}

	// 查找该值父节点
	parent := tree.Root.FindParent(value)

	// 1.删除根节点且根节点没有儿子,直接删除
	if parent == nil && node.Left == nil && node.Right == nil {
		// 置空后直接返回
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		// 2.删除节点有父节点，但没哟子树，直接删除

		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	} else if node.Left != nil && node.Right != nil {
		// 3.删除节点右两个子树，右子树比左子树大，用右子树中最小元素来替换删除的节点

		// 找右子树中最小的值
		minNode := node.Right
		for minNode != nil {
			minNode = minNode.Left
		}

		// 删除最小节点
		tree.Delete(minNode.Value)

		// 最小节点值替换被删除节点
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		// 4.删除的节点只有一个子树，直接替换被删除的节点即可

		// 父为空，表示删除的是根节点，替换树根
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}

		// 左子树不为空
		if node.Left != nil {
			// 如果删除的是父节点中的左节点，直接替换
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			// 右子树不为空
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}

// 中序遍历
// 先打印左子树，然后打印根节点的值，再打印右子树
func (tree *BinarySearchTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *BinarySearchTreeNode) MidOrder() {
	if node == nil {
		return
	}

	// 先打印左子树
	node.Left.MidOrder()

	// 按照子树打印根节点
	for i := 0; i < int(node.Times); i++ {
		fmt.Println(node.Value)
	}

	// 再打印右子树
	node.Right.MidOrder()
}
