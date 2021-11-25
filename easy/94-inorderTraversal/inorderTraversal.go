package _4_inorderTraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	current := root
	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, current.Val)
		current = current.Right
	}
	return res
}

func inorderTraversalMorris(root *TreeNode) []int {
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil {
			node := root.Left
			for node.Right != nil && node.Right != root {
				node = node.Right
			}
			if node.Right == nil {
				node.Right = root
				root = root.Left
				continue
			} else {
				res = append(res, root.Val)
				node.Right = nil
			}
		} else {
			res = append(res, root.Val)
		}
		root = root.Right
	}
	return res
}






func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}




func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return res
}


func morris(root *TreeNode) []int {
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil {
			node := root.Left
			for node.Right != nil && node.Right != root {
				node = node.Right
			}
			if node.Right == nil {
				node.Right = root
				root = root.Left
			} else {
				res = append(res, root.Val)
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