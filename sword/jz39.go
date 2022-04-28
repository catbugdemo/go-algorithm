package sword

// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

func majorityElement(nums []int) int {
	elem := make(map[int]int)
	Max := nums[0]
	for _, num := range nums {
		if _, ok := elem[num]; ok {
			elem[num]++
			// 判断是否是最大值
			if elem[Max] < elem[num]+1 {
				Max = num
			}
		} else {
			elem[num] = 1
		}
	}
	return Max
}

// 摩尔投票法
// 1.假设当前数为众数，遇到相同的数字则加一，否则减一
// 2.若前n个票为0(互相抵消),则设下一个数为当前数
// 3.重复1,2最后的当前数为众数
func majorityElement2(nums []int) int {
	x, votes := nums[0], 0
	for _, num := range nums {
		// 票数抵消
		if votes == 0 {
			x = num
		}
		if num == x {
			votes++
		} else {
			votes--
		}
	}
	return x
}
