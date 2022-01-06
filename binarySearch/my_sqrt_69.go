package binarySearch

import "math"

//实现int sqrt(int x)函数。
//
//计算并返回x的平方根，其中x 是非负整数。
//
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
//示例 1:
//
//输入: 4
//输出: 2
//

// mySqrt
func mySqrt(x int) int {
	return int(math.Floor(math.Sqrt(float64(x))))
}
