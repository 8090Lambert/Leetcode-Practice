package levelOrder

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		res = append(res, []int{})
		count := len(queue)
		for i := 0; i < count; i++ {
			node := queue[i]
			res[level] = append(res[level], node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = append([]*TreeNode{}, queue[count:]...)
		level++
	}
	return res
}

func backTrack(root *TreeNode, level int, res *[][]int) {
	if root != nil {
		if len(*res) == level {
			*res = append(*res, []int{})
		}
		(*res)[level] = append((*res)[level], root.Val)
		if root.Left != nil {
			backTrack(root.Left, level+1, res)
		}
		if root.Right != nil {
			backTrack(root.Right, level+1, res)
		}
	}
}