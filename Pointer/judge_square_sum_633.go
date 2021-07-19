package Pointer

import "math"

//给定一个非负整数c，你要判断是否存在两个整数 a 和 b，使得a^2 + b^2 = c 。
//
//
//
//示例 1：
//
//输入：c = 5
//输出：true
//解释：1 * 1 + 2 * 2 = 5

// judgeSquareSum 两数之和的变种
// 判断特殊数据
// 1.通过 来判断数据是否符合规则
// 2.判断特殊数据 他本身两个平方相加，或者他本身等于该结果
func judgeSquareSum(c int) bool {
	if c == 0{
		return true
	}

	hash := map[int]int{}
	//最大为c的平方根
	for i := 1; i <= int(math.Sqrt(float64(c))); i++ {
		if i*i*2 == c || i*i == c{
			return true
		}
		if _,ok := hash[c-i*i];ok{
			return true
		}
		//设置hash[i*i]不为空
		hash[i*i] = 1

	}
	return false
}

// judgeSquareSum2 优化
// 1.使用math.Sqrt
func judgeSquareSum2(c int) bool {
	for i := 0; i*i <= c ; i++ {
		//获取减去后开平方的数据
		rt := math.Sqrt(float64(c - i*i))
		//如果该数据是整数
		if rt == math.Floor(rt){
			return true
		}
	}
	return false
}