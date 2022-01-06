package greedy

import (
	"strings"
)

//字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
//
//输入：S = "ababcbacadefegdehijhklij"
//输出：[9,7,8]
//解释：
//划分结果为 "ababcbaca", "defegde", "hijhklij"。
//每个字母最多出现在一个片段中。
//像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

// partitionLabels
// 判断数组长度
// 1.对字母出现次数进行总结，包括：第一次出现位置，最后一次出现位置
// 2.根据第一次出现位置和最后一次出现位置作为范围
// 3.循环
// 4.如果范围重叠，取最小的第一次和最大的最后一次作为范围
func partitionLabels(s string) []int {
	//判断数组长度
	if len(s) == 0 {
		return []int{0}
	}
	if len(s) == 1 {
		return []int{1}
	}

	//总结
	//1.获取第一次出现次数和最后一次出现次数
	//2.将已经算好的字母转换为"_"
	//2.循环
	var ints [][]int
	tmpStr := s
	for i := 0; i < len(s); i++ {
		split := strings.Split(tmpStr, "")
		if split[i] == "_" {
			continue
		}
		index := strings.Index(tmpStr, split[i])
		lastIndex := strings.LastIndex(tmpStr, split[i])
		list := []int{index, lastIndex}
		ints = append(ints, list)
		//转换
		tmpStr = strings.ReplaceAll(tmpStr, split[i], "_")
	}

	//不需要对ints进行排序,因为ints是从最开始一个数据来的

	//循环，看范围是否重叠范围重叠, end > start
	var r []int
	start, end := ints[0][0], ints[0][1]
	for i := 1; i < len(ints); i++ {
		//[i][1]小于end，在原来的范围内
		if end > ints[i][1] {
			continue
		}
		//[i][1]大于end，但是[i][0]小于end,更新end
		if end > ints[i][0] {
			end = ints[i][1]
			continue
		}
		//没有end = ints[i][0] 的可能
		r = append(r, end-start+1)
		start = ints[i][0]
		end = ints[i][1]
	}
	//添加最后一个未获取的
	r = append(r, end-start+1)

	return r
}

func partitionLabels2(s string) []int {
	if len(s) == 0 {
		return []int{0}
	}
	if len(s) == 1 {
		return []int{1}
	}

	//获取最后一次的位置,不断循环会获取最后相同字母的最后一次出现时间
	m := make(map[int32]int)
	for k, v := range s {
		m[v] = k
	}

	//循环是否重叠
	result := []int{}
	start, end := 0, 0
	for k, v := range s {
		index := m[v]
		if end < index && end >= k {
			end = index
		} else if k > end {
			result = append(result, end-start+1)
			start = end + 1
			end = index
		}

	}
	result = append(result, end-start+1)
	return result
}
