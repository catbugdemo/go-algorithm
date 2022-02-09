package sword

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pre int整型一维数组
 * @param vin int整型一维数组
 * @return TreeNode类
 */

// 1.前序遍历的第一个一定根节点
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// 找到根节点
	if len(pre) == 0 {
		return nil
	}
	var index int
	for k, v := range vin {
		if pre[0] == v {
			index = k
			break
		}
	}
	root := &TreeNode{
		Val: pre[0],
	}

	root.Left = reConstructBinaryTree(pre[1:index+1], vin[0:index])
	root.Right = reConstructBinaryTree(pre[index+1:], vin[index+1:])
	return root
}
