package doubleNode

import "fmt"

type Node struct {
	Data interface{}
	pre  *Node
	next *Node
}

type DoubleList struct {
	headNode *Node
}

func (o *DoubleList) Add(data interface{}) {
	node := &Node{
		Data: data,
	}
	cur := o.headNode
	if cur == nil {
		o.headNode = node
	} else {
		for cur.next != nil {
			cur = cur.next
		}
		node.pre = cur
		cur.next = node
	}
}

func (o *DoubleList) Length() int {
	var count int
	cur := o.headNode
	if cur != nil {
		for cur != nil {
			cur = cur.next
			count++
		}
	}
	return count
}

func (o *DoubleList) RemoveNode(data interface{}) {
	cur := o.headNode
	if cur.Data == data {
		// 删除头结点
		newHead := cur.next
		newHead.pre = nil
		o.headNode = newHead
	} else {
		for cur != nil {
			if cur.Data == data {
				pre := cur.pre
				next := cur.next

				// 当前节点被跳过，会被当做垃圾回收掉
				next.pre = pre
				pre.next = next
			}
			cur = cur.next
		}
	}
}

// 从后向前
func (o *DoubleList) ShowFromTail() {
	tail := o.headNode
	for tail.next != nil {
		tail = tail.next
	}
	for tail.pre != nil {
		fmt.Printf("\t %v", tail.Data)
		tail = tail.pre
	}
}
