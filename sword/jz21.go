package sword

func exchange(nums []int) []int {
	result, one, two := []int{}, []int{}, []int{}
	for _, num := range nums {
		if num%2 == 1 {
			one = append(one, num)
		} else {
			two = append(two, num)
		}
	}
	result = append(result, one...)
	result = append(result, two...)
	return result
}
