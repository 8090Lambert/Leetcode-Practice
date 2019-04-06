package isSameTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p==nil && q==nil {
		return true
	}

	if p!=nil && q!=nil && p.Val==q.Val {
		return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
	}else {
		return false
	}
}
