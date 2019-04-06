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


func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}


	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree1(q.Left, p.Left) && isSameTree1(q.Right, p.Right)
	} else {
		return false
	}
}
