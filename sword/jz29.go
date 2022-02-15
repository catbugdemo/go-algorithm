package sword

func spiralOrder(matrix [][]int) []int {
	// 模拟行为，确定每个方向的边界
	if len(matrix) == 0 {
		return nil
	}
	var res []int
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		// 1.left -> right ,top 不变
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		// 左至右遍历到边界,top++ 往下挪一层
		top++
		// 2.top -> bottom
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		// 往左挪一层
		right--
		// 判断边界是否满足
		if top > bottom || left > right {
			break
		}
		// 3.right -> left
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		// 往上挪一层
		bottom--
		// 4.bottom -> top
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
