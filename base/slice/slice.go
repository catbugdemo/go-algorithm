package main

import (
	"fmt"
	"sync"
)

type Array struct {
	array []int
	len   int
	cap   int
	mu    sync.Mutex
}

func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}
	array := make([]int, cap, cap)
	s.array = array
	s.len = 0
	s.cap = cap
	return s
}

func (a *Array) Append(element int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// 大小等于容量，表示没有多余位置了
	if a.len == a.cap {
		newCap := 2 * a.len

		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]int, newCap, newCap)
		// 把老数据的数据移动到新数据
		for k, v := range a.array {
			newArray[k] = v
		}
		// 替换数组
		a.array = newArray
		a.cap = newCap
	}

	a.array[a.len] = element
	a.len += 1
}

// AppendMany  增加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

// Get 获取某个下标的元素
func (a *Array) Get(index int) int {
	// 越界了
	if a.len == 0 || a.len < index {
		panic("index over len")
	}
	return a.array[index]
}

// Len 返回真实长度
func (a *Array) Len() int {
	return a.len
}

// Cap 返回容量
func (a *Array) Cap() int {
	return a.cap
}

// 辅助打印
func Print(array *Array) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 第一个元素
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, array.Get(i))
			continue
		}
		result = fmt.Sprintf("%s%d", result, array.Get(i))
	}
	result = result + "]"
	return
}
