package sword

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	vide := new(ListNode)
	for pHead != nil {
		next := pHead.Next
		pHead.Next = vide
		vide = pHead
		pHead = next
	}
	return vide
}
