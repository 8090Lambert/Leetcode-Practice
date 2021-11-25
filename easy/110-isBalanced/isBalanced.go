package _10_isBalanced

import (
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	return isBalance(root) != -1
}

func isBalance (root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	left := isBalance(root.Left)
	right := isBalance(root.Right)
	if left != -1 && right != -1 && math.Abs(left-right) <= 1 {
		return math.Max(left, right) + 1
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

func isBalanced(root *TreeNode) bool {
	isBalance := true
	height(root, &isBalance)
	return isBalance
}

func height(node *TreeNode, isBalance *bool) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.Left, isBalance)
	if !(*isBalance) {
		return -1
	}
	rightHeight := height(node.Right, isBalance)
	if !(*isBalance) {
		return -1
	}

	if int(math.Abs(float64(leftHeight - rightHeight))) > 1 {
		*isBalance = false
		return -1
	}

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}



func isBalanced(root *TreeNode) bool {
	return h(root) >= 0
}

func h(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := h(root.Left)
	rightHeight := h(root.Right)
	if leftHeight == -1 || rightHeight == -1 || int(math.Abs(float64(leftHeight-rightHeight))) > 1 {
		return -1
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}