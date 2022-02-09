package sword

type CQueue struct {
	root *ListNode
	size int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	// 如果栈顶为空，增加节点
	if this.root == nil {
		this.root = new(ListNode)
		this.root.Val = value
	} else {
		// 否则新元素插入链表的末尾
		newNode := new(ListNode)
		newNode.Val = value

		// 找到末尾
		nowNode := this.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		// 添加节点
		nowNode.Next = newNode
	}
	this.size++
}

func (this *CQueue) DeleteHead() int {
	if this.size == 0 {
		return -1
	}

	// 出队
	topNode := this.root
	v := topNode.Val

	this.root = topNode.Next
	this.size--
	return v
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
