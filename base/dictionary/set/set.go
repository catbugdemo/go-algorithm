package main

import (
	"fmt"
	"sync"
	"unsafe"
)

// 实现不可重复集合 Set
// 集合指的是不可重复集合还是指统称的集合

type Set struct {
	m            map[int]struct{} // 用字典来实现，因为字段建不能重复
	len          int
	sync.RWMutex // 读写锁，实现并发安全
}

// 初始化一个集合
func NewSet(cap int64) *Set {
	return &Set{
		m: make(map[int]struct{}, cap),
	}
}

// 增加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()

	s.m[item] = struct{}{} // 实际王字典添加这个键
	s.len = len(s.m)       // 重新计算
}

// 删除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()

	// 集合没元素直接返回
	if s.len == 0 {
		return
	}
	delete(s.m, item) // 从字典删除这个键
	s.len = len(s.m)  // 重新计算元素数量
}

func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.m[item]
	return ok
}

// 查看集合大小
func (s *Set) Len() int {
	return s.len
}

// 是否为空
func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

// 清楚集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{} // 字典重新赋值
	s.len = 0                // 大小归零
}

// 将集合转化为列表
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// 空结构体的内存地址都一样，并不占用内存空间
func main() {
	a := struct{}{}
	b := struct {
	}{}
	if a == b {
		fmt.Printf("right:%p\n", &a)
	}
	fmt.Println(unsafe.Sizeof(a))
}
