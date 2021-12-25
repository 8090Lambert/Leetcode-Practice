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

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	backTrack1(root, targetSum, []int{}, &res)
	return res
}

func backTrack1(root *TreeNode, targetSum int, item []int, res *[][]int) {
	if targetSum < 0 || root == nil {
		return
	}
	item = append(item, root.Val)
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		*res = append(*res, item)
		return
	}
	backTrack1(root.Left, targetSum-root.Val, item, res)
	backTrack1(root.Right, targetSum-root.Val, item, res)
}