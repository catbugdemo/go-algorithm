package sword2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var res []int
	var check func(head *ListNode)
	check = func(head *ListNode) {
		if head == nil {
			return
		}
		check(head.Next)
		res = append(res, head.Val)
	}
	check(head)
	return res
}
