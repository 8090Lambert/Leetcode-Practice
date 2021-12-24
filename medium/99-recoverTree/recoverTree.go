package _9_recoverTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func recoverTree(root *TreeNode)  {
	stack := make([]*TreeNode, 0)
	var x, y, pre *TreeNode
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && root.Val < pre.Val {
			y = root
			if x == nil {
				x = pre
			} else {
				break
			}
		}
		pre = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}



func recoverTreeMorris(root *TreeNode)  {
	var pre, left, right *TreeNode
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
				if pre != nil && root.Val <= pre.Val {
					left = root
					if right == nil {
						right = pre
					}
				}
				pre = root
				node.Right = nil
				root = root.Right
			}
		} else {
			if pre != nil && root.Val <= pre.Val {
				left = root
				if right == nil {
					right = pre
				}
			}
			pre = root
			root = root.Right
		}
	}
	left.Val, right.Val = right.Val, left.Val
}