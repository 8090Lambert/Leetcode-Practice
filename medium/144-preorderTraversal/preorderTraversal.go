package _44_preorderTraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func PreorderTraversal (root *TreeNode) []int {
	res := make([]int, 0)
	//backTrack(root, &res)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, current.Val)
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	
	return res
}


func preorderTraversal(root *TreeNode) []int {

}

func backTrack (root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		if root.Left != nil {
			backTrack(root.Left, res)
		}
		if root.Right != nil {
			backTrack(root.Right, res)
		}
	}
}
