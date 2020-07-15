package _07_levelOrderBottom

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func LevelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0{
		tmp := make([]*TreeNode, 0)
		each := make([]int, 0)
		for _, val := range stack {
			if val.Left != nil {
				tmp = append(tmp, val.Left)
			}
			if val.Right != nil {
				tmp = append(tmp, val.Right)
			}
			each = append(each, val.Val)
		}
		stack = tmp
		result = append([][]int{each}, result...)
	}
	
	return result
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