package _13_pathSum


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSumDFS(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	backTrack(root, targetSum, 0, []int{}, &res)
	return res
}

func backTrack(root *TreeNode, target, current int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	current += root.Val
	path = append(path, current)
	if root.Left == nil && root.Right == nil && target == current {
		*res = append(*res, path)
		return
	}
	backTrack(root.Left, target, current, path, res)
	backTrack(root.Right, target, current, path, res)
	path = path[:len(path)-1]
}