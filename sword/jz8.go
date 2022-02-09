package sword

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

// 分情况
// 1.当前结点是父亲结点的左孩子
// 2.当前结点右孩子结点
// 3.前结点为父亲结点的右孩子结点
// 4.最尾结点
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return pNode
	}

	// 如果节点有右子树
	if pNode.Right != nil {
		pNode = pNode.Right
		for pNode.Left != nil {
			pNode = pNode.Left
		}
		return pNode
	}

	p := pNode.Next
	n := pNode
	for p != nil {
		if n == p.Right {
			n = p
			p = p.Next
			continue
		}
		return p
	}
	return nil
}
