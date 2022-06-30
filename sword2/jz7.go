package sword2

// 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 构建二叉树
// 使用递归进行创建
// 1. preorder 和 inorder 一定一样
// 2. 设置根节点
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 找到 中序遍历的根节点的所在位置
	var index int
	for i, num := range inorder {
		if num == preorder[0] {
			index = i
		}
	}
	root := &TreeNode{
		Val: preorder[0],
	}

	// 设置左右节点
	root.Left = buildTree(preorder[1:index+1], inorder[0:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}
