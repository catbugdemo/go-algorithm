package sword

func getKthFromEnd(head *ListNode, k int) *ListNode {
	nowNode := head
	for nowNode != nil {
		nowNode = nowNode.Next
		if k > 0 {
			k--
		} else {
			head = head.Next
		}
	}
	return head
}
