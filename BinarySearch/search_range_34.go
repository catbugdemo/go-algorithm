package BinarySearch

import "sort"

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回[-1, -1]。
//
//进阶：
//
//你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？
//
//
//示例 1：
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]

// searchRange 查询
// 进行二分查找
// 2.对数据进行二分查找
// 	分两次查找
// 	第一次查找第一次出现，第二次查找最后出现
func searchRange(nums []int, target int) []int {
	//查找第一次出现的情况
	left := sort.SearchInts(nums, target)

	//对left进行判断
	// left == len(nums) 意味着数据超出数组最大范围
	// nums[left] != target 意味着小于最小范围或者没找到
	// SearchInts在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	//通过查找比target大1的数据去 -1 可以查询到最后 target出现的次数
	right := sort.SearchInts(nums, target+1) - 1
	return []int{left, right}
}
