package sword

// 假定该数组是
func verifyPostorder(postorder []int) bool {
	var dfs func(l, r int) bool
	dfs = func(l, r int) bool {
		// 已经超出循环了
		if l > r {
			return true
		}
		root := postorder[r]
		k := l
		// 找到最右边
		for k < r && postorder[k] < root {
			k++
		}
		// 右子树都大于根节点
		for i := k; i < r; i++ {
			if root > postorder[i] {
				return false
			}
		}
		return dfs(l, k-1) && dfs(k, r-1)
	}
	return dfs(0, len(postorder)-1)
}
