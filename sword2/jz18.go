package sword2

// 删除链表节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	l := new(ListNode)
	l.Next = head

	// 遍历节点查找值
	point := l
	for point.Next != nil && point.Next.Val != val {
		point = point.Next
	}
	// 删除节点
	if point.Next != nil && point.Next.Val == val {
		point.Next = point.Next.Next
	}
	return l.Next
}
