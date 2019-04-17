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
