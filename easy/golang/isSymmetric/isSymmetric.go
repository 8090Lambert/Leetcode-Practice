package isSymmetric

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func IsDouble(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil && left.Val == right.Val{
		return IsDouble(left.Left, right.Right) && IsDouble(left.Right, right.Left)
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{root.Left, root.Right}
	for len(stack) > 0 {
		left := stack[0]
		right := stack[1]
		stack = stack[2:]
		if left == nil && right == nil {
			continue
		} else if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val  {
			return false
		}
		stack = append(stack, left.Right, right.Left, left.Left, right.Right)
	}
	return true
}