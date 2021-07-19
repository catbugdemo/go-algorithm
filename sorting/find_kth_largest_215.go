package sorting

import (
	"math/rand"
	"sort"
	"time"
)

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//示例 1:
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5

// findKthLargest
//
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 明天问下朱科伟组长
func findKthLargest2(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// 快速选择
// 不需要对其左右再进行排序，一般先打乱数组
func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	//结束条件
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q + 1, r, index)
	}
	return quickSelect(a, l, q - 1, index)
}

// randomPartition 随机选元法
// https://blog.csdn.net/ljianhui/article/deta
func randomPartition(a []int, l, r int) int {
	i := rand.Int() % (r - l + 1) + l
	// 随机获取一个a[i],将i换到最右边，作为选定中轴
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

// partition 切分操作
// 对于某个索引 j  nums[j]已经排定，即nums[j] 经过 partition 分切操作后会放置在它最终应该放置的地方
// nums[left] 到 nums[j-1] 中的所有元素都不大于 nums[j]
// nums[j+1] 到 nums[right] 中的所有元素都不小于 nums[j]
// 轴选择方法
func partition(a []int, l, r int) int {
	// x 为随机获取并确定的中轴
	x := a[r]
	//  l -1 防止数组越界，i 为左指针
	i := l - 1
	for j := l; j < r; j++ {
		// 从前向后扫描，j 为右指针
		// 放置将 j 放置在 i位置，
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	// i 之前的所有数据都小于 中轴
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

