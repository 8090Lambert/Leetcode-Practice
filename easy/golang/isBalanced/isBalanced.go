package isBalanced

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalance(root) != -1
}

func isBalance (root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := isBalance(root.Left)
	right := isBalance(root.Right)
	if left > 0 && right > 0 && math.Abs(float64(left - right)) <= 1{
		return int(math.Max(float64(left), float64(right))) + 1
	}
	return -1
}


// method 2
func isBalancedOptimize(root *TreeNode) bool {
	if root == nil {
		return true
	}

	depth := 0
	return optimize(root, &depth) != 0
}

func optimize(root *TreeNode, depth *int) int {
	if root == nil {
		*depth = 0
		return 1
	}
	var left, right int
	if optimize(root.Left, &left) != 0 && optimize(root.Right, &right) != 0 {
		gap := left-right
		if gap <= 1 && gap >= -1 {
			*depth = int(math.Max(float64(left), float64(right))) + 1
			return 1
		}
	}

	return 0
}
