package levelOrderBottom

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
