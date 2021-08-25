package _35_lowestCommonAncestor

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pVal := p.Val
	qVal := q.Val
	
	node := root
	for node != nil {
		current := node.Val
		if pVal > current && qVal > current {
			node = node.Right
		} else if pVal < current && qVal < current {
			node = node.Left
		} else {
			return node
		}
	}
	
	return nil
}

func lowestCommonAncestor1 (root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	pVal := p.Val
	qVal := q.Val
	current := root.Val
	
	if pVal > current && qVal > current {
		return lowestCommonAncestor1(root.Right, p, q)
	} else if pVal < current && qVal < current {
		return lowestCommonAncestor1(root.Left, p, q)
	} else {
		return root
	}
	
	return nil
}

