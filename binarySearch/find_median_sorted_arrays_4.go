package binarySearch

import "sort"

//给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的 中位数 。
//
//
//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2

// findMedianSortedArrays
// 未优化
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1) == 0 {
		return 0
	} else if len(nums1)%2 == 1 {
		return float64(nums1[len(nums1)/2])
	} else {
		return (float64(nums1[len(nums1)/2]) + float64(nums1[len(nums1)/2-1])) / 2
	}
	return 0
}
