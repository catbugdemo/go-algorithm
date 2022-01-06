package greedy

//给你一个长度为n的整数数组，请你判断在 最多 改变1 个元素的情况下，该数组能否变成一个非递减数列。
//
//我们是这样定义一个非递减数列的：对于数组中任意的i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

//输入: nums = [4,2,3]
//输出: true
//解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。

// checkPossibility
// 1.判断数组长度
//
func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			count++
			//波谷有两个以上
			if count > 1 {
				return false
			}
			//不考虑第一个数,当出现上面情况时，并且 nums[i-1] >nums[i+1],
			//num[i]是可删除的,判断nums[i+1]是否是可删除的
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}
