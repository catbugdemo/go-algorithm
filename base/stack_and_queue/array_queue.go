package stack_and_queue

import (
	"sync"
)

// 数组队列，先进先出
type ArrayQueue struct {
	array []string   // 底层切片
	size  int        // 队列元素数量
	lock  sync.Mutex // 并发安全锁
}

// 入队
func (queue *ArrayQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 放入切片,后进的元素放在数组最后面
	queue.array = append(queue.array, v)

	// 队中元素数量+1
	queue.size += 1
}

// 出队
func (queue *ArrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队列元素已空
	if queue.size == 0 {
		panic("empty")
	}

	// 队列最前面元素
	v := queue.array[0]

	/*
		直接原位移动，但缩容后继的空间不会被释放
		for i := 1; i< queue.size ;i++ {
			queue.array[i-1] = queue.array[i]
		}
		// 原数组缩容
		queue.array = queue.array[0 : queue.size-1]
	*/

	// 创建新的数组
	newArray := make([]string, queue.size-1, queue.size-1)
	copy(newArray, queue.array[:queue.size-1])
	queue.array = newArray

	// 队列元素中 -1
	queue.size -= 1
	return v
}
