package stack_and_queue

import "sync"

// 链表队列 LinkQueue
// 队列先进先出，和栈操作相反

type LinkQueue struct {
	root *LinkNode  // 链表起点
	size int        // 大小
	lock sync.Mutex // 并发安全锁
}

// 入队
func (queue *LinkQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 如果栈顶为空，增加节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		// 否则插入链表末尾
		newNode := new(LinkNode)
		newNode.Value = v

		// 一直遍历到链表尾部
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}

		// 添加尾链表
		nowNode.Next = newNode
	}

	// 队中元素数量+1
	queue.size += 1
}

// 出队
func (queue *LinkQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 边界条件
	if queue.size == 0 {
		panic("empty")
	}

	// 顶部元素出队列
	topNode := queue.root
	v := topNode.Value

	// 顶部元素后移
	queue.root = topNode.Next

	// 队中元素-1
	queue.size -= 1
	return v
}
