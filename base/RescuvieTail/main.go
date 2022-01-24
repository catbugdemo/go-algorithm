package main

import "fmt"

// 1*2*~*n

// 尾递归
func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}
	return RescuvieTail(n-1, a*n)
}

func main() {
	// 尾递归，第一个是尾 第二个是初始值
	fmt.Println(RescuvieTail(5, 1))
}
