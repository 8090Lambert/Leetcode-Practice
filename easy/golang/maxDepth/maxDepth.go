package maxDepth

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	} else {
		depth = 1
		return int(math.Max(float64(MaxDepth(root.Left)), float64(MaxDepth(root.Right)))) + depth
	}
}

// DFS
func maxDepth(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	depth := 0
	for len(stack) > 0 {
		depth += 1
		count := len(stack)
		for i := 0; i < count; i++ {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = append([]*TreeNode{}, stack[count:]...)
	}

	return depth
}
