package search

// 全排列
// 给定一个无重复数字的整数数组，求其所有的排列方式。

var res [][]int

func permute(nums []int) [][]int {
	backtrack(nums, len(nums), []int{})
	return res
}

// 回溯
// 递归重要的其实是对数的搜索
// 1. nums 原数组 2. 数组数据数量 3.需要添加的测试数据
func backtrack(nums []int, numsLen int, path []int) {
	// 如果数据为 0 , 则已到达叶子节点
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	for i := 0; i < numsLen; i++ {
		// 从 nums 中找到一个 数据
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backtrack(nums, len(nums), path)
		// 然后进行回溯 -- 下面都是对数据进行回溯
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		path = path[:len(path)-1]
	}
}
