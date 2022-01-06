package greedy

import (
	"sort"
)

//给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
//
//注意:
//
//可以认为区间的终点总是大于它的起点。
//区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
//

//输入: [ [1,2], [2,3], [3,4], [1,3] ]
//
//输出: 1
//
//解释: 移除 [1,3] 后，剩下的区间没有重叠。

// eraseOverlapIntervals 保留区间最少的，就能有更多权限留给其他
// 1.根据结束区间进行排序
// 2.进行比较，如果区间重叠，总数+1
func eraseOverlapIntervals(intervals [][]int) int {
	//判断是否为0,或1
	if len(intervals) == 0 || len(intervals) == 1 {
		return 0
	}
	//学习slice排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	//进行比较
	count, end := 0, intervals[0][1]
	//如果 开始区间小于 end意味着有重叠,否则更新end
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
			continue
		}
		end = intervals[i][1]
	}
	return count
}
