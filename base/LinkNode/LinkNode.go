package main

import "fmt"

/*
	链表由一个个数据节点组成，它是一个递归结构，要么它是空的，要么它存在指向另外一个数据节点的引用
*/

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	n1 := new(LinkNode)
	n1.Data = 1

	n2 := new(LinkNode)
	n2.Data = 2
	n1.NextNode = n2
	nowNode := n1
	for true {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			nowNode = nowNode.NextNode
			continue
		}
		break
	}
}
