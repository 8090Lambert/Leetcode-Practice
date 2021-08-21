package _26_invertTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func invertTreeWithLoop (root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	list := []*TreeNode{root}
	for len(list) != 0 {
		current := list[0]
		current.Left, current.Right = current.Right, current.Left
		list = append([]*TreeNode{}, list[1:]...)
		if current.Left != nil {
			list = append(list, current.Left)
		}
		if current.Right != nil {
			list = append(list, current.Right)
		}
	}
	
	return root
}
