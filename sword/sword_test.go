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
		obj := Constructors()
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

func TestJz15(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		weight := hammingWeight(11)
		assert.Equal(t, 3, weight)
	})
}

func TestJz16(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		pow := myPow(2, -2)
		assert.Equal(t, float64(1024), pow)
	})
}

func TestJz17(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		ints := printNumbers(1)
		fmt.Println(ints)
	})
}

func TestJz20(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		number := isNumber("1 ")
		fmt.Println(number)
	})
}

func TestJz23(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		top := new(ListNode)
		top.Val = 1

		l1 := top
		l2 := &ListNode{
			Val: 2,
		}
		l1.Next = l2

		l3 := &ListNode{
			Val: 3,
		}
		l2.Next = l3

		l4 := &ListNode{
			Val: 4,
		}
		l3.Next = l4

		l5 := &ListNode{
			Val: 5,
		}
		l4.Next = l5
		l5.Next = l3
		loop := EntryNodeOfLoop(top)
		fmt.Println(loop)
	})
}

func TestJz24(t *testing.T) {
	node := NewListNode([]int{1, 2, 3, 4, 5})
	list := reverseList(node)
	fmt.Println(list)
}

func TestJz25(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		l1 := NewListNode([]int{1, 2, 3})
		l2 := NewListNode([]int{1, 3, 4})
		mergeTwoLists(l1, l2)
	})

	t.Run("f", func(t *testing.T) {
		mergeTwoLists(nil, nil)
	})
}

func TestJz26(t *testing.T) {
	t.Run("t", func(t *testing.T) {
	})
}

func TestJz27(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		top := &TreeNode{
			Val: 4,
		}
		top.Left = &TreeNode{Val: 2}
		top.Right = &TreeNode{Val: 7}
		top.Left.Left = &TreeNode{Val: 1}
		top.Left.Right = &TreeNode{Val: 3}
		top.Right.Left = &TreeNode{Val: 6}
		top.Right.Right = &TreeNode{Val: 9}
		mirrorTree(top)

	})
}

func TestJz28(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		top := &TreeNode{Val: 1}
		top.Left = &TreeNode{Val: 2}
		top.Right = &TreeNode{Val: 2}
		top.Left.Right = &TreeNode{Val: 3}
		top.Right.Right = &TreeNode{Val: 3}
		check := isSymmetric(top)
		fmt.Println(check)
	})
	t.Run("f", func(t *testing.T) {
		symmetric := isSymmetric(nil)
		fmt.Println(symmetric)
	})
}

func TestJz31(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		check := validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})
		fmt.Println(check)
	})
	t.Run("f", func(t *testing.T) {
		check := validateStackSequences([]int{1, 2, 3, 0}, []int{2, 1, 3, 0})
		fmt.Println(check)
	})
}

func TestJz32(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		root := &TreeNode{Val: 3}
		root.Left = &TreeNode{Val: 9}
		root.Right = &TreeNode{Val: 20}
		root.Right.Left = &TreeNode{Val: 15}
		root.Right.Right = &TreeNode{Val: 7}
		order := levelOrder(root)
		fmt.Println(order)
	})
	t.Run("t", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		order := levelOrder(root)
		fmt.Println(order)
	})
}

func TestJz34(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		root := &TreeNode{Val: 5}
		root.Left = &TreeNode{Val: 4}
		root.Left.Left = &TreeNode{Val: 11}
		root.Left.Left.Left = &TreeNode{Val: 7}
		root.Left.Left.Right = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 8}
		root.Right.Left = &TreeNode{Val: 13}
		root.Right.Right = &TreeNode{Val: 4}
		root.Right.Right.Left = &TreeNode{Val: 5}
		root.Right.Right.Right = &TreeNode{Val: 1}
		sum := pathSum(root, 22)
		fmt.Println(sum)
	})
	t.Run("f", func(t *testing.T) {
		root := &TreeNode{Val: -2}
		root.Right = &TreeNode{Val: -3}
		sum := pathSum(root, -5)
		fmt.Println(sum)
	})
}
