package _03_zigzagLevelOrder

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func ZigzagLevelOrder (root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	level := 0
	for len(stack) > 0 {
		res = append(res, []int{})
		stackLen := len(stack)
		tmp := make([]*TreeNode, 0)
		for i := 0; i < stackLen; i++ {
			current := stack[i]
			res[level] = append(res[level], current.Val)
			if level & 1 == 0 {
				tmp = append([]*TreeNode{current.Left}, tmp...)
				tmp = append([]*TreeNode{current.Right}, tmp...)
			} else {
				tmp = append([]*TreeNode{current.Right}, tmp...)
				tmp = append([]*TreeNode{current.Left}, tmp...)
			}
		}
		stack = tmp
		level++
	}
	
	return res
}


func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	level := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		count := len(stack)
		tmp := []*TreeNode{}
		res = append(res, []int{})
		for i := 0; i < count; i++ {
			res[level] = append(res[level], stack[i].Val)
			if level & 1 == 0 {
				if stack[i].Left != nil {
					tmp = append([]*TreeNode{stack[i].Left}, tmp...)
				}
				if stack[i].Right != nil {
					tmp = append([]*TreeNode{stack[i].Right}, tmp...)
				}
			} else {
				if stack[i].Right != nil {
					tmp = append([]*TreeNode{stack[i].Right}, tmp...)
				}
				if stack[i].Left != nil {
					tmp = append([]*TreeNode{stack[i].Left}, tmp...)
				}
			}
		}
		level++
		stack = tmp
	}
	return res
}