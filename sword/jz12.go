package sword

// 使用辅助函数
type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

// 使用回溯法查找
func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	// 不重复查找
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		// 当前字符不匹配
		if board[i][j] != word[k] {
			return false
		}
		// 单词存在于网格中
		if k == len(word)-1 {
			return true
		}
		// 已查找
		vis[i][j] = true
		defer func() {
			// 回溯已访问单元格
			vis[i][j] = false
		}()
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
