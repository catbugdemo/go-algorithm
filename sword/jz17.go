package sword

import "math"

func printNumbers(n int) []int {
	m := int(math.Pow(10, float64(n)))
	result := make([]int, 0, m)
	for i := 1; i < m; i++ {
		result = append(result, i)
	}
	return result
}
