package pointer

type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle
// 使用hash将每次走过的数据存储起来
// 一旦遇到了此前遍历过的节点，就可以判定链表中存在环
func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return head
}
