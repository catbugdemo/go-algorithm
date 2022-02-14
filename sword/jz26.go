package sword

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 没有找到节点
	if A == nil || B == nil {
		return false
	}
	// 找到了节点判断是否相同
	if A.Val == B.Val {
		return recor(A, B)
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recor(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recor(A.Left, B.Left) && recor(A.Right, B.Right)
}
