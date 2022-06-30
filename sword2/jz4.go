package sword2

import "sort"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 边界值判断
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	// 找左下角数据
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// 方法二
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		index := sort.SearchInts(nums, target)
		if index < len(nums) && nums[index] == target {
			return true
		}
	}
	return false
}
