package sword

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 对称二叉树的中序遍历一定是对称的
// 用一个数组保存中序遍历，判断中序遍历是否一致
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return midOrder(root.Left, root.Right)
}

func midOrder(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return midOrder(left.Left, right.Right) && midOrder(left.Right, right.Left)
}
