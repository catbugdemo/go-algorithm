package main

import (
	"fmt"
	"sync"
)

// 定义
// 1.有节点间的层次关系，分为父节点和子节点
// 2.有唯一一个根节点，该根节点没有父节点
// 3.除了根节点，每个节点有且只有一个父节点
// 4.每一个节点本身以及它的后代也是一棵树，是一个递归的结构
// 5.没有后代的节点称为叶子节点，没有节点的树称为空树

// 二叉树： 每个节点最多只有两个儿子节点的树

// 满二叉树：叶子节点与叶子节点之间的高度差为 0 的二叉树，即整棵树是满的
// 数呈满三角形结构。在国外的定义，非叶子节点儿子儿子都是满的树就是满二叉树。我们以国内为准

// 完全二叉树：完全二叉树是由满二叉树而引出来的，设二叉树的深度为k，
// 除第k层外，其他各层的节点数都达到最大值，且第k层所有的节点都连续集中在最左边

// 实现
type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右子树
}

// 遍历二叉树
// 1.先序遍历：先访问根节点，再访问左子树，最后访问右子树
// 2.后序遍历：先访问左子树，再访问右子树，最后访问根节点
// 3.中序遍历：先访问左子树，再访问根节点，最后访问右子树
// 4.层次遍历：每一层从左到右访问每一个节点。

// 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印根节点
	fmt.Print(tree.Data, " ")
	// 再打印左子树
	PreOrder(tree.Left)
	// 最后打印右子树
	PreOrder(tree.Right)
}

// 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印左子树
	MidOrder(tree.Left)
	// 再打印根节点
	fmt.Print(tree.Data, " ")
	// 最后打印右子树
	MidOrder(tree.Right)
}

// 后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印左子树
	PostOrder(tree.Left)
	// 再打印右子树
	PostOrder(tree.Right)
	// 最后打印根节点
	fmt.Print(tree.Data, " ")
}

// 层次遍历较复杂，用到一种名叫广度遍历的方法，需要使用辅助的先进先出队列
func LayerOrder(treeNode *TreeNode) {
	if treeNode == nil {
		return
	}

	// 新建队列
	queue := new(LinkQueue)
	queue.Add(treeNode)
	for queue.size > 0 {
		// 不断出队
		element := queue.Remove()

		// 先打印节点值
		fmt.Print(element.Data, " ")

		// 左子树非空，入队列
		if element.Left != nil {
			queue.Add(element.Left)
		}

		// 右子树非空，入队列
		if element.Right != nil {
			queue.Add(element.Right)
		}
	}
}

type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value *TreeNode
}

func (queue *LinkQueue) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	newNode := &LinkNode{
		Value: v,
	}

	// 如果为空，增加节点
	if queue.root == nil {
		queue.root = newNode
	} else {
		// 不为空，放入链表尾部
		// 新节点
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		nowNode.Next = newNode
	}

	queue.size += 1
}

// 出队
func (queue *LinkQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队中元素已空
	if queue.size == 0 {
		panic("over limit")
	}

	topNode := queue.root
	v := topNode.Value

	// 出队
	queue.root = topNode.Next
	queue.size -= 1
	return v
}

func (queue *LinkQueue) Size() int {
	return queue.size
}

func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}

	fmt.Println("先序排序：")
	PreOrder(t)
	fmt.Println("\n中序排序：")
	MidOrder(t)
	fmt.Println("\n后序排序")
	PostOrder(t)
	fmt.Println("\n层次遍历：")
	LayerOrder(t)
}
