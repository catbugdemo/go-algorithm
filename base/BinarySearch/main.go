package main

import "fmt"

// 二分查找递归算法
// 1 5 9 15 81 89 123 189 333
// https://goa.lenggirl.com/#/basic/rescuvie?id=%e4%b8%89%e3%80%81%e4%be%8b%e5%ad%90%ef%bc%9a%e4%ba%8c%e5%88%86%e6%9f%a5%e6%89%be

// 理论上，所有递归都能转换为非递归，但是递归代码可读性更高
// 该数据应该已经排好
func B(array []int, target int) int {
	var binarySearch func(l, r int) int
	binarySearch = func(l, r int) int {
		if l >= r {
			// 出界了 找不到
			return -1
		}

		// 从中间开始查找
		mid := (r-l)>>1 + l
		if array[mid] == target {
			return target
		} else if array[mid] > target {
			// 中间数比把目标大 ，从左边找
			return binarySearch(l, mid-1)
		} else {
			// 中间数比目标数小，从右边找
			return binarySearch(mid+1, r)
		}
	}
	return binarySearch(0, len(array))
}

func main() {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 500
	result := B(array, target)
	fmt.Println(target, result)

	target = 189
	result = B(array, target)
	fmt.Println(target, result)
}
