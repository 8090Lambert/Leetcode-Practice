package _8_isValidBST

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	min := math.MinInt64
	for len(stack) != 0 || root != nil {
		for root != nil {
			root = root.Left
			if root != nil {
				stack = append(stack, root)
			}
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= min {
			return false
		}
		min = root.Val
		root = root.Right
	}
	return true
}


func isValidBSTMorris(root *TreeNode) bool {
	min := math.MinInt64
	for root != nil {
		if root.Left != nil {
			node := root.Left
			for node.Right != nil && node.Right != root {
				node = node.Right
			}
			if node.Right == nil {
				node.Right = root
				root = root.Left
			} else {
				if root.Val <= min {
					return false
				}
				min = root.Val
				node.Right = nil
				root = root.Right
			}
		} else {
			if root.Val <= min {
				return false
			}
			min = root.Val
			root = root.Right
		}
	}
	return true
}
