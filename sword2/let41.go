package sword2

import "fmt"

// 1.存取在 字典中
// 2.遍历查找
func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 将复数改成 n+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	// 做标记
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			fmt.Println(num - 1)
			nums[num-1] = -abs(nums[num-1])
		}
	}
	// 查询最小数据
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
