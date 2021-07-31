package Pointer

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

func rotate(nums []int, k int)  {
	copy(nums,append(nums[len(nums)-k:],nums[:len(nums)-k]...))
}

