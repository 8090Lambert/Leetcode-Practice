package _00_isSameTree

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

func isSameTreeLoop(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	stackp := []*TreeNode{p}
	stackq := []*TreeNode{q}
	for len(stackp) > 0 {
		p := stackp[0]
		q := stackq[0]
		stackp = stackp[1:]
		stackq = stackq[1:]
		if !check(p, q) || !check(p.Left, q.Left) || !check(p.Right, q.Right) {
			return false
		}
		if p != nil {
			if p.Left != nil {
				stackp = append(stackp, p.Left)
				stackq = append(stackq, q.Left)
			}
			if p.Right != nil {
				stackp = append(stackp, p.Right)
				stackq = append(stackq, q.Right)
			}
		}
	}
	return true
}

func check (p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return true
}


func isSameTreeOther(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && q.Val == p.Val {
		return isSameTreeOther(q.Left, p.Left) && isSameTreeOther(q.Right, p.Right)
	} else {
		return false
	}
}