package Pointer

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//示例 1：
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"

// minWindow 判断最短子串出现 t
// 1.先判断特殊情况
// 2.将t中所有的数据转为utf-8编码转换
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	//将t中所有的数据转为utf-8编码转换
	//统计t中的所有数据
	hash := make([]int, 256)
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}
	//l:为作指针指向，count:t的总数,max：，results结果
	l, count, max, results := 0, len(t), len(s)+1, ""

	//r：为右指针指向
	for r := 0; r < len(s); r++ {
		//找到当前s[r]在 hash中的数据 -1
		hash[s[r]]--

		//说明该s[r]存在于t中
		if hash[s[r]] >= 0 {
			count--
		}

		//l为左指针指向，当出现s[l:r]重复出现同一t内数据时
		//移动左指针直到再次满足条件
		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}

		//如果t的所有值都判断到了，判断最长长度
		//该条件判断了所有的符合t的所有字段
		if count == 0 && max > r-l+1 {
			max = r - l + 1
			results = s[l : r+1]
		}
	}
	return results
}
