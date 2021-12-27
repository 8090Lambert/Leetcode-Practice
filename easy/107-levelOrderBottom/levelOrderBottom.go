package _07_levelOrderBottom

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		levels := make([]*TreeNode, 0)
		count := len(stack)
		item := make([]int, 0)
		for i := 0; i < count; i++ {
			item = append(item, stack[i].Val)
			if stack[i].Left != nil {
				levels = append(levels, stack[i].Left)
			}
			if stack[i].Right != nil {
				levels = append(levels, stack[i].Right)
			}
		}
		stack = levels
		res = append([][]int{item}, res...)
	}
	return res
}

func levelOrderBottomOther(root *TreeNode) [][]int {
	res := make([][]int, 0)
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) != 0 {
		count := len(stack)
		item := make([]int, 0, count)
		for i := 0; i < count; i++ {
			cur := stack[i]
			item = append(item, cur.Val)
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}
		res = append(res, item)
		stack = stack[count:]
	}
	start, end := 0, len(res)-1
	for start < end {
		res[start], res[end] = res[end], res[start]
		start++
		end--
	}
	return res
}