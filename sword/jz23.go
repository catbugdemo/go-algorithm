package sword

// 快慢指针解决问题
// 即 2 *(a +b) = a + n *(b + c)
// a = (n-1) (b + c) + c
// 可认为 a = c

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	// 无环
	if pHead == nil || pHead.Next == nil {
		return nil
	}
	// 设置快慢指针
	slow, fast := pHead, pHead
	for fast != nil {
		if slow == fast {
			// 即快慢指针相遇，重置 fast，当再次相遇即为圆环点
			fast = pHead
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return fast
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}

// 笨，但是简单
func EntryNodeOfLoop2(pHead *ListNode) *ListNode {
	m := map[*ListNode]struct{}{}
	for pHead != nil {
		if _, ok := m[pHead]; ok {
			return pHead
		}
		m[pHead] = struct{}{}
		pHead = pHead.Next
	}
	return nil
}
