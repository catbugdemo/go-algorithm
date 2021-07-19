package Pointer

//给定一个字符串 s ，找出至多包含 k 个不同字符的最长子串 T。
//
//示例 1:
//
//输入: s = "eceba", k = 2
//输出: 3
//解释: 则 T 为 "ece"，所以长度为 3。

// lengthOfLongestSubstringKDistinct
// 可根据题76进行求解
// 用滑动窗口算法进行求解
// 同时用hash来求解
// 1.右窗口滑动,将数据存储
// 2.判断存储的数据是否符合标准
// 3.不符合存储数据移动左窗口
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	max := 0
	hash := map[uint8]int{}
	for l,r := 0,0; r < len(s); r++ {
		hash[s[r]]+=1
		if len(hash) <= k && max < r-l+1 {
			max =r-l+1
		}else {
			for len(hash) > k {
				hash[s[l]]-=1
				if hash[s[l]] == 0{
					delete(hash,s[l])
				}
				l++
			}
		}
	}
	return max
}
