package sword2

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 找出最小值
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}
