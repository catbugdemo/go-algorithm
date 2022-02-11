package sword

func deleteNode(head *ListNode, val int) *ListNode {
	topNode := head
	var parent *ListNode
	for topNode != nil {
		if topNode.Val != val {
			parent = topNode
			topNode = topNode.Next
			continue
		}

		// 需要删除的数据
		if topNode == head {
			head = head.Next
		} else {
			parent.Next = topNode.Next
		}
		break
	}
	return head
}
