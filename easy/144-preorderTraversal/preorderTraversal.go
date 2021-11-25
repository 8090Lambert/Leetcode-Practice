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


func preorderTraversalOther(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1].Right
		stack = append([]*TreeNode{}, stack[:len(stack)-1]...)
	}
	return res
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

func preorderTraversalMorris(root *TreeNode) []int {
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil {
			node := root.Left
			for node.Right != nil && node.Right != root {
				node = node.Right
			}
			if node.Right == nil {
				res = append(res, root.Val)
				node.Right = root
				root = root.Left
			} else {
				node.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}

func preorderTraversalWithMorris(root *TreeNode) []int {
	res := make([]int, 0)
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				res = append(res, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			} else {
				p2.Right = nil
			}
		} else {
			res = append(res, p1.Val)
		}
	}
	return res
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}












