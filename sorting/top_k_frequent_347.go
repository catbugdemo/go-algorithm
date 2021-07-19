package sorting

import "sort"

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
//
//示例 1:
//
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]

// topKFrequent 使用桶排序
// 1.先对数据 通过map 进行分组
// 2.降序排序
// 3.返回相应的值
// map 本身就是一个桶实现
// 桶排序通过 使用 map 必须通过一个 数组记录 map的key 否则无法进行排序
func topKFrequent(nums []int, k int) []int {
	// 分组
	mList := make(map[int]int)
	// 记录key值
	s := make([]int, 0)

	for _, num := range nums {
		if _, ok := mList[num]; ok {
			mList[num]++
		} else {
			mList[num] = 1
			s = append(s,num)
		}
	}
	// 将记录 key 根据 mList[s[i]] > mList[s[j]] 排序
	sort.Slice(s, func(i, j int) bool {
		return mList[s[i]] > mList[s[j]]
	})
	return s[:k]
}
