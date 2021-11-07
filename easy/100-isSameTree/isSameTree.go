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

func isSameTreeDfS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	stack := []*TreeNode{p, q}
	for len(stack) != 0 {
		count := len(stack)
		l, r := stack[count-1], stack[count-2]
		stack = stack[:count-2]
		if !compare(l, r) || !compare(l.Left, r.Left) || !compare(l.Right, r.Right) {
			return false
		}
		if l.Left != nil {
			stack = append(stack, l.Left, r.Left)
		}
		if r.Right != nil {
			stack = append(stack, l.Right, r.Right)
		}
	}
	return len(stack) == 0
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
		if !compare(p, q) || !compare(p.Left, q.Left) || !compare(p.Right, q.Right) {
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

func compare(p, q *TreeNode) bool {
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