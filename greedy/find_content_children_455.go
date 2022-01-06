package greedy

import (
	"sort"
)

//输入: g = [1,2,3], s = [1,1]
//输出: 1
//解释:
//你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
//虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
//所以你应该输出1。

// findContentChildren
// 1.对 孩子(g) 和 饼干(s) 进行从小到大排序
// 2.遍历s, 当 g[0] 满足向后递增 s从被满足开始
// 当index >len : index out of range
// 先要用是否为最小的饼，满足最小的小孩、
// 可以 去掉 count :减少用时
func findContentChildren(g []int, s []int) int {
	// 排序 g 和 s
	sort.Ints(g)
	sort.Ints(s)

	//遍历
	count, indexg, indexs := 0, 0, 0
	for {
		if indexg < len(g) && indexs < len(s) {
			if g[indexg] <= s[indexs] {
				indexg++
				count++
			}
			indexs++
		} else {
			break
		}
	}

	return count
}

// findContentChildren2 简化了 count 计数，同时将indexs的自动增加添加到了 for循环中
// 再次简化
func findContentChildren2(g []int, s []int) int {
	if !sort.IntsAreSorted(g) { //去掉该判断会减少消耗时间，但会增加内存消耗
		sort.Ints(g)
	}
	if !sort.IntsAreSorted(s) {
		sort.Ints(s)
	}

	indexg := 0
	for i := 0; i < len(s); i++ {
		if indexg < len(g) {
			if g[indexg] <= s[i] {
				indexg++
			}
		} else {
			break
		}
	}
	return indexg
}
