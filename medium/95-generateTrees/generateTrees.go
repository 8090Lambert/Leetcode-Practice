package _5_generateTrees

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper (start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTree := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTree := helper(start, i-1)
		rightTree := helper(i+1, end)
		for _, left := range leftTree {
			for _, right := range rightTree {
				curr := &TreeNode{i, nil, nil}
				curr.Left = left
				curr.Right = right
				allTree = append(allTree, curr)
			}
		}
	}
	return allTree
}
