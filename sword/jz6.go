package sword

type ListNode struct {
	Val  int
	Next *ListNode
}

var printList []int

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	if head == nil {
		return nil
	}
	printListFromTailToHead(head.Next)
	printList = append(printList, head.Val)
	return printList
}
