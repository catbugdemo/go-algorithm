package trivial

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

// twoSum 两数之和
// 1.循环 第一个数和后面的数依次相加
// 2.如果等于跳出循环
// 时间复杂度 O(N^2)
// 空间复杂度 O(N)
func twoSum(nums []int, target int) []int {
	ints := make([]int, 0, 10)
	for i := 0; i < len(nums); i++ {
		for j := 1 + i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ints = append(ints, i)
				ints = append(ints, j)
				break
			}
		}
	}
	return ints
}

// twoSum2 利用hash的存储key值判断是否存在
// 1.循环 : 判断 与 target - value 的数据是否存在 ：
// 	存在输出，不存在 continue
// 2.记录数据
// 时间复杂度 O(N)
// 空间复杂度 O(N)
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {  // 永远 k<i ,因为 i 是根据循环递增， k 是之前存储的 i
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}
