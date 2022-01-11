package search

// 给定一个整数 n 和一个整数 k，求在 1 到 n 中选取 k 个数字的所有组合方法。

var cres [][]int

// 这道题和 46 题相似，都使用回溯法进行计算
func combine(n int, k int) [][]int {
	p := make([]int, 0, k)
	cbacktrack(1, n, k, p)
	return cres
}

// 回溯法
// 1.创建树
// 2.找到如果 数组的长度=k,则意味着可以判断
func cbacktrack(start, end, k int, path []int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		cres = append(cres, tmp)
		return
	}

	// 使用回溯
	for i := start; i <= end; i++ {
		path = append(path, i)
		cbacktrack(i+1, end, k, path)
		path = path[:len(path)-1]
	}
}
