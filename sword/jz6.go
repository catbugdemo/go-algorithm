package sword

type ListNode struct {
	Val  int
	Next *ListNode
}

var result []int

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	if head == nil {
		return nil
	}
	printListFromTailToHead(head.Next)
	result = append(result, head.Val)
	return result
}
