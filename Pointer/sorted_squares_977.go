package Pointer

import "sort"

// 有序数组的平方

// 最简单
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}

// 双指针算法
// 该方法通过左右两指针进行缩进
func sortedSquares2(nums []int) []int {
	var ret = make([]int, len(nums))
	start, end := 0, len(nums)-1
	index := len(nums) - 1
	for start <= end {
		if nums[start]*nums[start] > nums[end]*nums[end] {
			ret[index] = nums[start] * nums[start]
			start++
		} else {
			ret[index] = nums[end] * nums[end]
			end--
		}
		index--
	}
	return ret
}
