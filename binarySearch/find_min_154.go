package binarySearch

import "sort"

// findMin
// 进行排序的算法(根本没有二分查找)
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	return nums[0]
}

//没有用二分查找，而是直接通过每个数据进行查询
func findMin2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	return min
}
