package main

import "fmt"

// 递归
// 但是因为使用了 位运算符，每次重复的调用都使得运算的链条不断加长，系统不得不使用栈进行数据保存和恢复。

func main() {
	r := Rescuvie(5)
	fmt.Println(r)
}

func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}
	return n * Rescuvie(n-1)
}
