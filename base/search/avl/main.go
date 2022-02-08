package main

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
