package main

// 2-3 一颗严格自平衡的多路查找树
// 是一颗三叉树

// 左倾红黑树,定义
// 1.根节点的链接是黑色的
// 2.红链接均为左链接
// 3.没有任何一个节点同时和两条红链接向量
// 4.任意一个节点到叶子结点的所有路径，经过的黑链接数量相同
// 也就是该树是完美黑色平衡，
// 由于红链接都在左边，所以这种红黑树又称左倾红黑树。左倾红黑树与 2-3 树一一对应，只要将左链接画平

// 定义颜色

const (
	RED   = true
	BLACK = false
)

// 左倾红黑树
type LLRBTree struct {
	Root *LLRBTNode
}

// 左倾红黑树节点
type LLRBTNode struct {
	Value int64      // 值
	Times int64      // 值出现的次数
	Left  *LLRBTNode // 左子树
	Right *LLRBTNode // 右子树
	Color bool       // 父亲指向该节点的链接颜色
}

// 节点的颜色
func IsRed(node *LLRBTNode) bool {
	if node == nil {
		return false
	}
	return node.Color == RED
}

// 左旋
func RotateLeft(h *LLRBTNode) *LLRBTNode {
	if h == nil {
		return nil
	}

	x := h.Right
	h.Right = x.Left
	x.Left = h
	x.Color = h.Color
	h.Color = RED
	return x
}

// 右旋
func RotateRight(h *LLRBTNode) *LLRBTNode {
	if h == nil {
		return nil
	}

	x := h.Left
	h.Left = x.Right
	x.Left = h
	x.Color = h.Color
	h.Color = RED
	return x
}

// 颜色转换
func ColorChange(h *LLRBTNode) {
	if h == nil {
		return
	}
	h.Color = !h.Color
	h.Left.Color = !h.Left.Color
	h.Right.Color = !h.Right.Color
}

// 左倾红黑树添加元素
func (tree *LLRBTree) Add(value int64) {
	// 根节点开始添加元素
	tree.Root = tree.Root.Add(value)
	// 根节点的链接永远都是黑色的
	tree.Root.Color = BLACK
}

func (node *LLRBTNode) Add(value int64) *LLRBTNode {
	// 插入的节点为空,将其链接设置为红色，返回
	if node == nil {
		return &LLRBTNode{
			Value: value,
			Color: RED,
		}
	}

	// 插入的元素重复
	if value == node.Value {
		node.Times++
	} else if value > node.Value {
		// 插入的元素比节点值大
		node.Right = node.Right.Add(value)
	} else {
		node.Left = node.Left.Add(value)
	}

	// 辅助变量
	nowNode := node

	// 右链接为红色，那么进行左旋，确保是左倾
	if IsRed(nowNode.Right) && !IsRed(nowNode.Left) {
		nowNode = RotateLeft(nowNode)
	} else {
		// 连续两个左链接为红色，那么进行右旋
		if IsRed(nowNode.Left) && IsRed(node.Left.Left) {
			nowNode = RotateRight(nowNode)
		}
		// 旋转后，可能左右链接都为红色，需要变色
		if IsRed(nowNode.Left) && IsRed(nowNode.Right) {
			ColorChange(nowNode)
		}
	}
	return nowNode
}

// 删除
