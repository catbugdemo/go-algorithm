package sword

// 即找到最大的能够到达的位置
// 使用回溯法
// 同时使用辅助函数
// 最后计算为大方格 - 小方格
func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var check func(i, j, count int) int
	check = func(i, j, count int) int {
		if !IfEnter(i, j, k) || vis[i][j] {
			return count
		}

		count++
		vis[i][j] = true
		// 继续向4个方向移动
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < m && 0 <= newJ && newJ < n && !vis[newI][newJ] {
				count = check(newI, newJ, count)
			}
		}
		return count
	}
	res := check(0, 0, 0)
	return res
}

// 判断能否进入方格
func IfEnter(i, j, k int) bool {
	var tmp int
	for i > 0 {
		tmp += i % 10
		i /= 10
	}
	for j > 0 {
		tmp += j % 10
		j /= 10
	}
	if tmp <= k {
		return true
	} else {
		return false
	}
}
