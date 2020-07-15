package _04_maxDepth

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

type Node struct {
	Depth int
	Val *TreeNode
}
func maxDepthDFS(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	stack := make([]Node, 0)
	stack = append(stack, Node{Depth:1,Val:root})
	depth = 0
	for len(stack) > 0 {
		current := stack[0]
		stack = append([]Node{}, stack[1:]...)
		if current.Val != nil {
			depth = int(math.Max(float64(depth), float64(current.Depth)))
			stack = append(stack, Node{Depth:depth+1, Val:current.Val.Left})
			stack = append(stack, Node{Depth:depth+1, Val:current.Val.Right})
		}
	}
	return depth
}

// BFS
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
