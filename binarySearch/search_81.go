package binarySearch

import "sort"

//已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
//
//在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
//
//给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。
//
//
//
//示例1：
//
//输入：nums = [2,5,6,0,0,1,2], target = 0
//输出：true

//search
func search(nums []int, target int) bool {
	sort.Ints(nums)
	result := sort.SearchInts(nums, target)
	if result == len(nums) || nums[result] != target {
		return false
	}
	return true
}

//search2 对代码进行优化,不进行数组的重新排序
func search2(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}
	return false
}
