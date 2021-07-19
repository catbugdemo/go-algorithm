package Pointer

import (
	"sort"
)

//给你两个有序整数数组nums1 和 nums2，请你将 nums2 合并到nums1中，使 nums1 成为一个有序数组。
//
//初始化nums1 和 nums2 的元素数量分别为m 和 n 。你可以假设nums1 的空间大小等于m + n，这样它就有足够的空间保存来自 nums2 的元素。

//示例 1：
//
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]

// merge
// 根据m截取 nums1 ,根据 n截取 nums2
// 拼接两数组
// 排序
// 输出
func merge(nums1 []int, m int, nums2 []int, n int)  {
	ints := append(nums1[:m], nums2[:n]...)
	sort.Ints(ints)
}

//优化,时间不变，减少内存消耗，有个疑问，前一个容量不够怎么办
func merge2(nums1 []int, m int, nums2 []int, _ int) {
	copy(nums1[m:],nums2)
	sort.Ints(nums1)
}