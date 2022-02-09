package sword

// 使用二分法查找
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		pivot := left + (right-left)/2
		if numbers[pivot] < numbers[right] {
			right = pivot
		} else if numbers[pivot] > numbers[right] {
			left = pivot + 1
		} else {
			right--
		}
	}
	return numbers[left]
}
