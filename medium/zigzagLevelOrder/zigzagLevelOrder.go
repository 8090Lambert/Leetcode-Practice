package zigzagLevelOrder

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder (root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	level := 0
	for len(stack) > 0 {
		res = append(res, []int{})
		stackLen := len(stack)
		tmp := make([]*TreeNode, 0)
		for i := 0; i < stackLen; i++ {
			current := stack[i]
			res[level] = append(res[level], current.Val)
			if level & 1 == 0 {
				tmp = append([]*TreeNode{current.Left}, tmp...)
				tmp = append([]*TreeNode{current.Right}, tmp...)
			} else {
				tmp = append([]*TreeNode{current.Right}, tmp...)
				tmp = append([]*TreeNode{current.Left}, tmp...)
			}
		}
		stack = tmp
		level++
	}
	
	return res
}
