package BinarySearch

//给定一个只包含整数的有序数组，每个元素都会出现两次，唯有一个数只会出现一次，找出这个数。
//
//示例 1:
//
//输入: [1,1,2,3,3,4,4,8,8]
//输出: 2

// singleNonDuplicate (未优化)
// 从零开始，每二一跳
// 如果当前数等于之后的数据
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] == nums[i+1] {
			continue
		}
		return nums[i]
	}
	return nums[len(nums)-1]
}
