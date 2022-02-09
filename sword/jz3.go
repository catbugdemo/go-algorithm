package sword

func duplicate(numbers []int) int {
	// write code here
	m := make(map[int]int)
	max := -1
	for _, num := range numbers {
		// 不存在该数据
		if _, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num]++
			max = num
		}
	}
	return max
}
