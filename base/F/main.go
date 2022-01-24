package main

import "fmt"

// 尾运算
// 1 1 2 3 5 8 13 21 ... N-1 N 2N-1
// b 永远作为结果进行返回

func F(end int, a int, b int) int {
	if end == 0 {
		return a
	}
	return F(end-1, b, a+b)
}

func main() {
	fmt.Println(F(5, 1, 1))
}
