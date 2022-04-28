package sword

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cacheNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cacheNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cacheNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

// 分为两次复制
func copyRandomList(head *Node) *Node {
	cacheNode = map[*Node]*Node{}
	return deepCopy(head)
}
