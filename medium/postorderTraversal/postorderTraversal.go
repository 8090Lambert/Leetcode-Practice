package postorderTraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal (root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{current.Val}, res...)
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}
	
	return res
}

func backTrack (root *TreeNode, res *[]int) {
	if root != nil {
		if root.Left != nil {
			backTrack(root.Left, res)
		}
		if root.Right != nil {
			backTrack(root.Right, res)
		}
		*res = append(*res, root.Val)
	}
}
