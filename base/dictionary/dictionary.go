package main

import "fmt"

// 字典是存储键值对的数据结构，把一个键和一个值映射起来
// 映射，键不能重复。在某些教程中，这种结构体可能成为符号表，关联数据或映射

func main() {
	// 新建一个容量为4的字典 map
	m := make(map[string]int, 4)

	// 放宰割键值对
	m["dog"] = 1
	m["cat"] = 2
	m["hen"] = 3

	fmt.Println(m)

	// 查找 hen
	which := "hen"
	if v, ok := m[which]; ok {
		// 找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		// 找不到
		fmt.Println("not find:", which)
	}

	// 查到 ccc
	which = "ccc"
	if v, ok := m[which]; ok {
		// 找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		// 没找到
		fmt.Println("not find:", which)
	}
}

// 字典的实现有两种方式
// 哈希表 HashTable 和红黑树 RBTree
// Golang 语言中字典 map 的实现由哈希表实现，具体可参考标准库 runtime 下的 map.go 文件
