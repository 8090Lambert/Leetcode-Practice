package _11_minDepth

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 1
	} else {
		if root.Left != nil && root.Right != nil {
			return int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right)))) + 1
		} else if root.Left != nil {
			return minDepth(root.Left) + 1
		} else if root.Right != nil {
			return minDepth(root.Right) + 1
		} else {
			return 1
		}
	}
}

func minDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else if root.Right != nil && root.Left == nil {
		return minDepth(root.Right) + 1
	} else {
		return int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right)))) + 1
	}
}


func minDepthBFS(root *TreeNode) int {
	depth := 0
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) != 0 {
		depth += 1
		count := len(stack)
		for i := 0; i < count; i++ {
			if stack[i].Left == nil && stack[i].Right == nil {
				return depth
			}
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[count:]
	}
	return depth
}