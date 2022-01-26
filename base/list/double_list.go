package list

import "sync"

// 列表 List: 可以存放数据的数据结构，数据按顺序排序，可以一次入队和出队，
// 列表有需要关系，可以取出某个序号的数据。先进先出的 队列(queue) 和先进后出的栈(stack) 都是列表
// 大家也经常听说一种叫 线性表的数据结构体，表示具有相同特性数据元素的有限序列

// 双端列表 -- 缓存数据库 Redis 的列表List 基本类型就是用它来实现的

// DoublieList 双端列表，双端队列
type DoubleList struct {
	head *ListNode  // 指向链表头部
	tail *ListNode  // 指向链表尾部
	len  int        // 列表长度
	lock sync.Mutex // 为了进度并发安全 pop 弹出操作
}

type ListNode struct {
	pre   *ListNode // 前驱节点
	next  *ListNode // 后驱节点
	value string    // 值
}

// 列表节点普通操作
func (node *ListNode) GetValue() string {
	return node.value
}

// 获取节点前驱
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// 获取节点后驱
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// 是否存在前驱节点
func (node *ListNode) HasPre() bool {
	return node.pre != nil
}

// 是否存在后驱节点
func (node *ListNode) HasNext() bool {
	return node.next != nil
}

// 是否为空节点
func (node *ListNode) IsNill() bool {
	return node.pre != nil
}

// 返回列表长度
func (list *DoubleList) Len() int {
	return list.len
}

// 从头部开始某个位置前插入新节点

// AddNodeFromHead 从头部开始，添加节点到第 N+1 个元素之前
// N=0表示添加到第一个元素之前，表示新节点成为新的头部
// N=1表示添加到第二个元素之前，以此类推
func (list *DoubleList) AddNodeFromHead(n int, v string) {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 如果索引查过或等于列表长度，一定找不到，直接 panic
	if n != 0 && n >= list.len {
		panic("index out")
	}

	// 获取头部，往后遍历拿到第 N+1 个位置的元素
	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 新节点
	newNode := new(ListNode)
	newNode.value = v

	// 如果定位到的节点为空，表示列表为空，将新节点设置为新头部和新尾部
	if node.IsNill() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到的节点，他的前驱
		pre := node.pre

		// 如果定位到的节点前驱为 nil,那么定位到的节点为链表头部，需要换头部
		if pre.IsNill() {
			newNode.next = node
			node.pre = newNode
			// 新节点成为头节点
			list.head = newNode
		} else {
			// 将新节点插入到定位到的节点前
			pre.next = newNode
			newNode.pre = pre

			// 定位到的节点的后驱节点 node.next 现在链接到新的节点上
			node.next.pre = newNode
			newNode.next = node.next
		}
	}
	// 列表长度 +1
	list.len += 1
}

// 从尾部开始某个位置后插入新节点

// AddNodeFromTail 从尾部开始，添加点到第N+1个元素之后，
//N=0 表示添加到第一个元素之后，表示新节点成为新的尾部，
//M=1 表示添加到第二个元素之后
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	if n != 0 && n >= list.len {
		panic("index out")
	}

	// 往前遍历拿到第 N+1 个位置的元素
	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	// 新节点
	newNode := new(ListNode)
	newNode.value = v

	// 如果定位到的节点为空，表示列表为空
	if node.IsNill() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到的节点
		next := node.next

		// 如果定位到的节点后驱nil，那么定位到的节点为链表尾部，需要换尾部
		if next.IsNill() {
			node.next = newNode
			newNode.pre = node

			// 新节点成为尾部
			list.tail = newNode
		} else {
			// 将新节点插入到定位到的节点之后
			newNode.pre = node
			node.next = newNode

			// 定位到的节点的后驱节点链接在新节点之后
			newNode.next = next
			next.pre = newNode
		}
	}

	// 列表长度 +1
	list.len = list.len + 1
}

// IndexFromHead 从头部开始往后找，获取第 N+1 个位置的节点，索引从0开始
func (list *DoubleList) IndexFromHead(n int) *ListNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	// 获取头结点
	node := list.head
	// 往后遍历
	for i := 1; i <= n; i++ {
		node = node.next
	}
	return node
}

// IndexFromTail 从尾部开始往前找
func (list *DoubleList) IndexFromTail(n int) *ListNode {
	if n >= list.len {
		return nil
	}

	// 获取尾部节点
	node := list.tail
	// 往前遍历
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	return nil
}

// 从头部开始移除并返回某个位置节点
func (list *DoubleList) PopFromHead(n int) *ListNode {
	// 加锁
	list.lock.Lock()
	defer list.lock.Unlock()

	if n >= list.len {
		return nil
	}

	// 获取头部
	node := list.head

	// 往后遍历
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 移除的节点的前驱和后驱
	pre := node.pre
	next := node.next
	if pre.IsNill() && next.IsNill() {
		// 唯一节点
		list.head = nil
		list.tail = nil
	} else if pre.IsNill() {
		// 移除的头结点
		list.head = next
		next.pre = nil
	} else {
		// 移除的是中间节点
		pre.next = next
		next.pre = pre
	}

	// 节点 -1
	list.len = list.len - 1
	return node
}
