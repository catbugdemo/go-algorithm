package sword2

// 使用二分法查询相应最少数据
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	// 通过二分法
	for left < right {
		mid := left + (right-left)>>1
		// 如果 mid 大于 right 意味着最小值在右边
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			// 小于 ，意味着 最小值在左边
			right = mid
		} else {
			// 等于要减
			right--
		}
	}
	return numbers[left]
}
