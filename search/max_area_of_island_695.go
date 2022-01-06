package search

// 给定一个二维的 0-1 矩阵，其中 0 表示海洋，1 表示陆地。单独的或相邻的陆地可以形成岛
// 屿，每个格子只与其上下左右四个格子相邻。求最大的岛屿面积。

// 使用深度搜索
// 1.当出现 1 (能够搜索的时候使用递归进行搜索)
func maxAreaOfIsland(grid [][]int) int {
	area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area = max(dfs(grid, i, j), area)
			}
		}
	}
	return area
}

var (
	// (i+1,j) (i-1,j) (i,j) (i,j+1) (i,j-1)
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

// 深度搜索 (创建数组原因：对于四个方向的遍历，可以创造辅函数，每相邻两位即为上下左右四个方向之一)
func dfs(grid [][]int, i, j int) (res int) {
	if grid[i][j] == 1 {
		grid[i][j] = 0
		res++
		for k := 0; k < 4; k++ {
			mx, my := i+dx[k], j+dy[k]
			if mx >= 0 && mx <= len(grid) && my >= 0 && my <= len(grid[0]) {
				res += dfs(grid, i, j)
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
