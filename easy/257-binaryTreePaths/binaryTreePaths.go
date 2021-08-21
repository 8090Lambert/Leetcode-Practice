package _57_binaryTreePaths

import "strconv"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func binaryTreePathsDFS(root *TreeNode) []string {
	paths := make([]string, 0)
	findPath(root, "", &paths)
	return paths
}

func findPath(root *TreeNode, path string, paths *[]string) {
	if root != nil {
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			*paths = append(*paths, path)
		} else {
			path += "->"
			findPath(root.Left, path, paths)
			findPath(root.Right, path, paths)
		}
	}
}

func binaryTreePathsBFS(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	paths := []string{}
	stack := []*TreeNode{root}
	list := []string{strconv.Itoa(root.Val)}
	for len(stack) != 0 {
		count := len(stack)
		for i := 0; i < count; i++ {
			cur := stack[i]
			base := list[i]
			if cur.Left == nil && cur.Right == nil {
				paths = append(paths, base)
				continue
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
				list = append(list, base + "->" + strconv.Itoa(cur.Left.Val))
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				list = append(list, base + "->" + strconv.Itoa(cur.Right.Val))
			}
		}
		stack = stack[count:]
		list = list[count:]
	}
	return paths
}
