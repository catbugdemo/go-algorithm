package sword2

// 创建一个字典去存取其中相应的内容
// 如果字典存在则自动存取其中的值
func findRepeatNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		} else {
			m[num] = 0
		}
	}
	return 0
}
