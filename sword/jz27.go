package sword

// 从节点开始转换
// 使用递归
func mirrorTree(root *TreeNode) *TreeNode {
	// 根节点为空
	if root == nil {
		return nil
	}
	// 为叶子节点
	if root.Left == nil && root.Right == nil {
		return root
	}
	// 左右节点相互交换
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}
