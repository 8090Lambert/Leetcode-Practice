package plusOne

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	tree := BuildTree(nums, 0, len(nums) - 1)
	return tree
}

func BuildTree (nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	
	mid := left + (right - left) / 2
	tree := &TreeNode{
		Val:nums[mid],
	}
	tree.Left = BuildTree(nums, left, mid - 1)
	tree.Right = BuildTree(nums, mid + 1, right)
	
	return tree
}
