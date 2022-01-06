package greedy

import (
	"sort"
)

//给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。
//
//输入：points = [[10,16],[2,8],[1,6],[7,12]]
//输出：2
//解释：对于该样例，x = 6 可以射爆 [2,8],[1,6] 两个气球，以及 x = 11 射爆另外两个气球

//	findMinArrowShots	找出最少的箭的个数
// 	1.数组长度判断(边界值判断 len==0 or len==1)
//	2.对数组进行排序
//	3.如果有重合,缩小重合区域
//	4.如果没有重合,从新重合开始计算
//	5.计数
func findMinArrowShots(points [][]int) int {
	//数组判断
	if len(points) == 0 {
		return 0
	}
	if len(points) == 1 {
		return 1
	}

	//对数组进行排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	//循环进行比较,从1开始，至少需要一支箭
	count, end := 1, points[0][1]
	for i := 1; i < len(points); i++ {
		//重合,缩小重合
		if points[i][0] > end {
			count++
			end = points[i][1]
		}
	}

	return count
}
