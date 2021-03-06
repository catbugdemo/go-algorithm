package pointer

//给定一个头结点为 head 的非空单链表，返回链表的中间结点。
//
//如果有两个中间结点，则返回第二个中间结点。
//
//
//
//示例 1：
//
//输入：[1,2,3,4,5]
//输出：此列表中的结点 3 (序列化形式：[3,4,5])
//返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
//注意，我们返回了一个 ListNode 类型的对象 ans，这样：
//ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// middleNode 返回此中间节点
// 两次循环
func middleNode(head *ListNode) *ListNode {
	tmp := head
	var index int
	for tmp.Next != nil {
		index++
		tmp = tmp.Next
	}

	// 第二次循环
	if index%2 == 0 {
		index = index / 2
	} else {
		index = index/2 - 1
	}
	tmp = head
	for i := 0; i < index; i++ {
		tmp = tmp.Next
	}
	return tmp
}

// 优化方法 快慢指针
// 快指针走两步，慢指针走一步
func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

//
