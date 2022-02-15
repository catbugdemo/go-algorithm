package sword

func pathSum(root *TreeNode, target int) [][]int {
	var result [][]int
	var check func(root *TreeNode, add []int, count int)
	check = func(root *TreeNode, add []int, count int) {
		if root == nil {
			return
		}
		add = append(add, root.Val)
		count += root.Val
		if count == target && root.Left == nil && root.Right == nil {
			fin := make([]int, len(add))
			copy(fin, add)
			result = append(result, fin)
			return
		}
		// 这个时候小于，需要继续递归
		check(root.Left, add, count)
		check(root.Right, add, count)
	}
	check(root, []int{}, 0)
	return result
}
