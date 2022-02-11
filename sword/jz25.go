package sword

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	top := new(ListNode)
	node := top
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return top.Next
}

// 递归调用
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var node *ListNode
	if l1.Val < l2.Val {
		node = l1
		node.Next = mergeTwoLists2(l1.Next, l2)
	} else {
		node = l2
		node.Next = mergeTwoLists2(l1, l2.Next)
	}
	return node
}
