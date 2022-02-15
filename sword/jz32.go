package sword

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var reslevel [][]int
	var helper func(root *TreeNode, index int)
	helper = func(root *TreeNode, index int) {
		if root == nil {
			return
		}
		if len(reslevel) == index {
			reslevel = append(reslevel, []int{})
		}
		reslevel[index] = append(reslevel[index], root.Val)
		helper(root.Left, index+1)
		helper(root.Right, index+1)
	}
	helper(root, 0)
	return reslevel
}
