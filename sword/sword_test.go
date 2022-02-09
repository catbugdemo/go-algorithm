package sword

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewListNode(list []int) *ListNode {
	topNode := new(ListNode)
	tmpNode := topNode
	for _, v := range list {
		tmpNode.Next = &ListNode{
			Val: v,
		}
		tmpNode = tmpNode.Next
	}
	return topNode.Next
}

func TestPrintListFromTailToHead(t *testing.T) {
	node := NewListNode([]int{1, 2, 3})
	head := printListFromTailToHead(node)
	fmt.Println(head)
}

func TestReConstructBinaryTree(t *testing.T) {
	pre := []int{1, 2, 3}
	vin := []int{2, 1, 3}
	tree := reConstructBinaryTree(pre, vin)
	PreOrder(tree)
}

func TestGetNext(t *testing.T) {
	root := &TreeLinkNode{Val: 8}
	root.Left = &TreeLinkNode{Val: 6}
	root.Right = &TreeLinkNode{Val: 10}
	root.Left.Next = root
	root.Right.Next = root
	root.Left.Left = &TreeLinkNode{Val: 5}
	root.Left.Right = &TreeLinkNode{Val: 7}
	root.Left.Left.Next = root.Left
	root.Left.Right.Next = root.Left
	root.Right.Left = &TreeLinkNode{Val: 9}
	root.Right.Right = &TreeLinkNode{Val: 11}
	root.Right.Left.Next = root.Right
	root.Right.Right.Next = root.Right

	//r5 := &TreeLinkNode{Val: 5}
	next := GetNext(root)
	fmt.Println(next)
}

// 前序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

func TestJz3(t *testing.T) {
	i := duplicate([]int{2, 3, 1, 0, 2, 5, 3})
	fmt.Println(i)
}

func TestJz4(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		find := Find(7, [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}})
		fmt.Println(find)
	})
	t.Run("false", func(t *testing.T) {
		find := Find(16, nil)
		fmt.Println(find)
	})
}

func TestJz5(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		space := replaceSpace("We Are Happy")
		fmt.Println(space)
	})
}

func TestJz9(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		obj := Constructor()
		obj.AppendTail(3)
		head := obj.DeleteHead()
		deleteHead := obj.DeleteHead()
		fmt.Println(head)
		fmt.Println(deleteHead)
	})
}

func TestJz10(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		i := fib(43)
		fmt.Println(i)
	})
}

func TestJz10_2(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		ways := numWays(7)
		fmt.Println(ways)
	})
}

func TestJz11(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		array := minArray([]int{1, 3, 5})
		assert.Equal(t, 1, array)
	})
}

func TestJz13(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		count := movingCount(2, 3, 1)
		assert.Equal(t, 3, count)
	})
}
