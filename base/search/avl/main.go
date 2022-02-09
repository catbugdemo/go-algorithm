package main

import (
	"fmt"
)

// avl 数 -- 严格自平衡的二叉查找树
// 1.首先它是一颗二叉查找树
// 2.任意一个节点的左右子树最大高度差为 1

// 平衡二叉查找树比较那一理解的是添加和删除元素时的调整操作

// AVL树
type AVLTree struct {
	Root *AVLTreeNode // 树根节点
}

// AVL 节点
type AVLTreeNode struct {
	Value  int64 // 值
	Times  int64 // 值出现的次数
	Height int64 // 该节点作为树根节点，树的高度，方便一算平衡因子
	Left   *AVLTreeNode
	Right  *AVLTreeNode
}

// 更新树的高度
func (node *AVLTreeNode) UpdateHeight() {
	if node == nil {
		return
	}
	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	// 哪个子树高算哪棵的
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	// 高度加上自己那一层
	node.Height = maxHeight + 1
}

// 计算树的平衡因子，也就是左右子树的高度差
func (node *AVLTreeNode) BalanceFactor() int64 {
	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return leftHeight - rightHeight
}

// 添加元素 -- 旋转
// 4种情况
// 1.在右子树上插上右儿子导致失衡，左旋，转一次
// 2.在左子树上插上左儿子导致失衡，右旋，转一次
// 3.在左子树上插上右儿子导致失衡，先左后右旋，转两次
// 4.在右子树上插上右儿子导致失衡，先右后左旋，转两次

// 左子树插左儿子：单右旋
func RightRotation(Root *AVLTreeNode) *AVLTreeNode {
	// 只有 Pivot和B, root 位置改变
	Pivot := Root.Left
	B := Pivot.Right
	Pivot.Right = Root
	Root.Left = B

	// 只有Root 和 Pivot 变化了高度
	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

// 右子树插右儿子：单左旋
func LeftRotation(Root *AVLTreeNode) *AVLTreeNode {
	Pivot := Root.Right
	B := Pivot.Left
	Pivot.Left = Root
	Root.Right = B

	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

// 右子树插左儿子：先右后左旋
func RightLeftRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Right = RightRotation(node.Right)
	return LeftRotation(node)
}

// 左子树插右儿子：先左后
func LeftRightRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Left = LeftRotation(node.Left)
	return RightRotation(node)
}

// 添加元素
func (tree *AVLTree) Add(value int64) {
	// 往树根添加元素，会返回新的树根
	tree.Root = tree.Root.Add(value)
}

func (node *AVLTreeNode) Add(value int64) *AVLTreeNode {
	// 添加值到根节点 node ，如果 node 为空，那么让值成为新的根节点，树的高度为1
	if node == nil {
		return &AVLTreeNode{Value: value}
	}
	// 如果值重复，什么都不用做，直接更新次数
	if node.Value == value {
		node.Times = node.Times + 1
		return node
	}

	// 辅助变量
	var newTreeNode *AVLTreeNode

	if value > node.Value {
		// 插入的值大于节点值，要从右子树继续插入
		node.Right = node.Right.Add(value)
		// 平衡因子，插入右子树后，要确保树根左子树的高度不能比右子树低一层
		factor := node.BalanceFactor()
		// 右子树高度变高了，导致左子树 - 右子树的高度从 -1 变成了 -2
		if factor == -2 {
			if value > node.Right.Value {
				// 表示在右子树上插上右儿子导致失衡，需要单左旋
				newTreeNode = LeftRotation(node)
			} else {
				// 表示在右子树上插上左儿子导致失衡，先右后左旋：
				newTreeNode = RightLeftRotation(node)
			}
		}
	} else {
		// 插入值小于节点值，要从左子树继续插入
		node.Left = node.Left.Add(value)
		// 平衡因子
		factor := node.BalanceFactor()
		if factor == 2 {
			if value < node.Left.Value {
				// 右旋
				newTreeNode = RightRotation(node)
			} else {
				// 先左后右旋
				newTreeNode = LeftRightRotation(node)
			}
		}
	}

	if newTreeNode == nil {
		// 表示什么旋转都没有，根节点没变
		node.UpdateHeight()
		return node
	} else {
		// 旋转了，更新树的高度
		newTreeNode.UpdateHeight()
		return newTreeNode
	}
}

// 找出最小值节点
func (tree *AVLTree) FindMinValue() *AVLTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *AVLTreeNode) FindMinValue() *AVLTreeNode {
	if node.Left == nil {
		// 左子树为空表示已经找到了最小值
		return nil
	}
	// 一直左子树递归
	return node.Left.FindMinValue()
}

// 找出最大值节点
func (tree *AVLTree) FindMaxValue() *AVLTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *AVLTreeNode) FindMaxValue() *AVLTreeNode {
	if node.Left == nil {
		// 右子树为空表示已经找到了最大值
		return nil
	}
	// 一直右子树递归
	return node.Right.FindMaxValue()
}

// 查找指定节点
func (tree *AVLTree) Find(value int64) *AVLTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.Find(value)
}

func (node *AVLTreeNode) Find(value int64) *AVLTreeNode {
	if value == node.Value {
		//找到该节点
		return node
	} else if value < node.Value {
		// 小于该节点从左子树开始找
		if node.Left == nil {
			return node
		}
		return node.Left.Find(value)
	} else {
		// 从右子树开始找
		if node.Right == nil {
			return node
		}
		return node.Right.Find(value)
	}
}

// 中序遍历
func (tree *AVLTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *AVLTreeNode) MidOrder() {
	if node == nil {
		return
	}
	// 先打印左子树
	node.Left.MidOrder()

	// 打印根节点
	for i := 0; i < int(node.Times); i++ {
		fmt.Println(node.Value)
	}

	// 再打印右子树
	node.Right.MidOrder()
}

// 删除 4 种情况
// 1.

func (node *AVLTreeNode) Delete(value int64) *AVLTreeNode {
	if node == nil {
		// 如果是空树，直接返回
		return nil
	}
	if value < node.Value {
		// 从左子树开始删除
		node.Left = node.Left.Delete(value)
		// 删除后更新该子树高度
		node.Left.UpdateHeight()
	} else if value > node.Value {
		node.Right = node.Right.Delete(value)
		// 更新子树高度
		node.Right.UpdateHeight()
	} else {
		// 找到了该值对应节点
		// 该节点没有左右子树
		// 情况1，删除节点没有儿子，直接删除
		if node.Left == nil && node.Right == nil {
			// 直接返回 nil ,表示直接该值删除
			return nil
		}

		// 情况2，删除节点小有两个子树，选高度更高的子树下的节点替换
		if node.Left != nil && node.Right != nil {
			// 左子树更高
			if node.Left.Height > node.Right.Height {
				maxNode := node.Left
				for maxNode.Right != nil {
					maxNode = maxNode.Right
				}

				// 最大节替换被删除节点
				node.Value = maxNode.Value
				node.Times = maxNode.Times

				// 删除最大节点
				node.Left = node.Left.Delete(maxNode.Value)
				// 删除后更新该子树高度
				node.Left.UpdateHeight()
			} else {
				// 右子树更高
				minNode := node.Right
				for minNode.Left != nil {
					minNode = minNode.Left
				}

				// 最小值替换被删除节点
				node.Value = minNode.Value
				node.Times = minNode.Times

				// 把最小节点删除
				node.Right = node.Right.Delete(minNode.Value)
				node.Right.UpdateHeight()
			}
		} else {
			// 只有一个子树
			if node.Left != nil {
				// 情况3，删除节点只有左子树，因为树的特性，左子树只有一个节点
				node.Value = node.Left.Value
				node.Times = node.Left.Times
				node.Height = 1
				node.Left = nil
			} else if node.Right != nil {
				node.Value = node.Right.Value
				node.Times = node.Right.Times
				node.Height = 1
				node.Right = nil
			}
		}
		// 找到值后进行替换并返回
		return node
	}

	// 左右子数递归删除节点后需要平衡
	var newNode *AVLTreeNode
	// 删除了右子树节点，左边变比右边高了，不平衡
	if node.BalanceFactor() == 2 {
		if node.Left.BalanceFactor() >= 0 {
			newNode = RightRotation(node)
		} else {
			newNode = LeftRightRotation(node)
		}
	} else if node.BalanceFactor() == -2 {
		//删除了左子树节点，右边比左边搞了，不平衡
		if node.Right.BalanceFactor() <= 0 {
			newNode = LeftRotation(node)
		} else {
			newNode = RightLeftRotation(node)
		}
	}

	if newNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newNode.UpdateHeight()
		return newNode
	}
}
