package main

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
	"math"
	"sync"
)

// 实现拉链 hash 表

const (
	// 扩容因子
	expandFactor = 0.75
)

// 哈希表
type HashMap struct {
	array        []*keyPairs // 哈希表数组，每个元素时一个键值对
	capacity     int         // 数组容量
	len          int         // 已添加键值对元素数量
	capacityMask int         // 掩码，等于 capacity -1
	// 增删键值对时，需要考虑并发安全
	lock sync.Mutex
}

// 键值对，连成一个链表
type keyPairs struct {
	key   string      // 键
	value interface{} // 值
	next  *keyPairs   // 下一个键值对
}

// 初始化哈希表
func NewHashMap(capacity int) *HashMap {
	// 默认大小为 16
	defaultCapacity := 1 << 4
	if capacity <= defaultCapacity {
		// 如果传入的大小小于默认大小，那么使用默认大小 16
		capacity = defaultCapacity
	} else {
		// 否则，实际大小为大于 capacity 的第一个 2^k
		capacity = 1 << (int(math.Ceil(math.Log2(float64(capacity)))))
	}

	// 新建一个 哈希表
	m := new(HashMap)
	m.array = make([]*keyPairs, capacity, capacity)
	m.capacity = capacity
	m.capacityMask = capacity - 1
	return m
}

// 返回哈希表已添加元素数量
func (m *HashMap) Len() int {
	return m.len
}

// 求 key 的哈希值
var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

func (m *HashMap) hashIndex(key string, mask int) int {
	// 求哈希
	hash := hashAlgorithm([]byte(key))
	// 求下边
	index := hash & uint64(mask)
	return int(index)
}

// 哈希表添加键值对
func (m *HashMap) Put(key string, value interface{}) {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()

	// 键值对要放的哈希表数组下标
	index := m.hashIndex(key, m.capacityMask)
	// 哈希表数组下标的元素
	element := m.array[index]

	// 元素为空，表示空链表，没有哈希冲突，直接赋值
	if element == nil {
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		// 链表最后一个键值对
		var lastPairs *keyPairs

		// 遍历链表查看元素是否存在，存在则替换值，否则找到最后一个键值对
		for element != nil {
			// 键值对存在，那么更新值并返回
			if element.key == key {
				element.value = value
				return
			}

			lastPairs = element
			element = element.next
		}

		// 找不到键值对，将新键值对添加到链表尾端
		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}
	}

	// 新的哈希表数量
	newLen := m.len + 1

	// 如果超出扩容因子，需要扩容
	if float64(newLen)/float64(m.capacity) >= expandFactor {
		// 新建一个原来两倍大小的哈希表
		newM := new(HashMap)
		newM.array = make([]*keyPairs, 2*m.capacity, 2*m.capacity)
		newM.capacity = 2 * m.capacity
		newM.capacityMask = 2*m.capacity - 1

		// 遍历老的哈希表，将键值对重新哈希到新的哈希表中
		for _, pairs := range m.array {
			for pairs != nil {
				// 直接递归Put
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		// 替换老的哈希表
		m.array = newM.array
		m.capacity = newM.capacity
		m.capacityMask = newM.capacityMask
	}

	m.len = newLen
}

func (m *HashMap) Range() {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, pairs := range m.array {
		for pairs != nil {
			fmt.Printf("'%v'='%v'", pairs.key, pairs.value)
			pairs = pairs.next
		}
	}
	fmt.Println()
}

// 总结
// 哈希表查找，是一种用空间换时间的查找算法，时间复杂度能达到 O(1)
// 最坏情况下退化到查找链表:O(n)，
// 但均匀性很好的哈希算法以及合适空间大小的数组，在很大概率避免了最坏情况

// 哈希表在添加元素时会进行伸缩，会造成较大的性能消耗，所以有时候会用到其他的查找算法
// 树查找算法
