package stack_and_queue

import "sync"

// 链表栈 -- 链表形式的下压栈，后进先出
type LinkStack struct {
	root *LinkNode  // 链表起点
	size int        // 栈的元素数量
	lock sync.Mutex // 并发安全使用的锁
}

// LinkNode 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}

// 入栈
func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 如果栈顶为空，那么增加节点
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		// 否则新元素插入链表的头部
		// 原来的链表
		preNode := stack.root

		// 配置新节点
		newNode := new(LinkNode)
		newNode.Value = v

		// 连接节点
		newNode.Next = preNode

		// 新节点作为头部
		stack.root = newNode
	}

	// 栈中元素数量+1
	stack.size += 1
}

// 出栈
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈中元素为空
	if stack.size == 1 {
		panic("empty")
	}

	// 顶部元素出栈
	topNode := stack.root
	v := topNode.Value

	// 连接
	stack.root = topNode.Next
	// 栈中元素数量-1
	stack.size -= 1
	return v
}

// 获取栈顶元素
func (stack *LinkStack) Peek() string {
	// 栈中元素为空
	if stack.size == 0 {
		panic("empty")
	}

	// 顶部元素值
	v := stack.root.Value
	return v
}

// 判断大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// 是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}
