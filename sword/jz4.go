package sword

// 1.在一维中如果该树大于当前，从下一个数组中查找
func Find(target int, array [][]int) bool {
	var re bool
	if array == nil {
		return re
	}
	// write code here
	var i, j int
	length := len(array)
	for target >= array[i][j] {
		var count int
		// 找到了
		if target == array[i][j] {
			re = true
			break
		}
		// 分别和 i+1 和 j+1 比较
		if length > i+1 && target >= array[i+1][j] {
			i++
		} else {
			count++
		}
		if length > j+1 && target >= array[i][j+1] {
			j++
		} else {
			count++
		}
		if count == 2 {
			break
		}
	}
	return re
}
