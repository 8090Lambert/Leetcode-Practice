package inorderTraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal (root *TreeNode) []int {
	res := []int{}
	backTrack(root, &res)
	return res
}

func backTrack (root *TreeNode, res *[]int) {
	if root != nil {
		if root.Left != nil {
			*res = append(*res, root.Left.Val)
		}
		*res = append(*res, root.Val)
		if root.Right != nil {
			*res = append(*res, root.Right.Val)
		}
	}
}

func R (root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		res = append(res, current.Val)
		current = current.Right
	}
	
	return res
}
