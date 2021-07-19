package Pointer

//输入：numbers = [2,7,11,15], target = 9
//输出：[1,2]
//解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

// twoSum
// 通过map存储 number[index] 和 index
func twoSum(numbers []int, target int) []int {
	m := map[int]int{}
	for k, v := range numbers {
		// index 永远小于 k, 因为index 是之前的k
		if index,ok := m[target-v];ok{
			return []int{index+1,k+1}
		}
		m[v] = k
	}
	return nil
}