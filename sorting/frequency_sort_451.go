package sorting

import (
	"sort"
	"strings"
)

// 给定一个字符串，请将字符串里的字符按照出现的频率降序排列。

/*
输入:
"tree"

输出:
"eert"

解释:
'e'出现两次，'r'和't'都只出现一次。
因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。

*/

// frequencySort 使用桶排序
// 先对每个字符进行计算，排序，然后根据大小进行输出
// 对map 进行排序需要一个额外的数组存储数据
func frequencySort(s string) string {
	// 分组
	v := make(map[string]int,len(s))
	// 记录key值
	k := make([]string,0)
	// 计算每个字符共有多少个
	for _, str := range strings.Split(s,"") {
		if _,ok := v[str];ok {
			v[str]++
			continue
		}
		v[str] = 1
		k = append(k,str)
	}

	// 将记录 key 根据 v[k[i]] > v[k[j]] 排序
	sort.Slice(k, func(i, j int) bool {
		return v[k[i]] > v[k[j]]
	})

	// 重新打印
	var builder strings.Builder
	for _, index := range k {
		for i := 0; i <v[index]; i++ {
			builder.WriteString(index)
		}
	}
	return builder.String()
}