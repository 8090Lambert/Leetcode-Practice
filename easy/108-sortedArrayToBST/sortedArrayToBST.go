package _08_sortedArrayToBST

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

func Build(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end - start) >> 1
	root := &TreeNode{Val:nums[mid]}
	root.Left = Build(nums, start, mid-1)
	root.Right = Build(nums, mid+1, end)
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	stack := make([]int, 0, len(nums))
	nodes := make([]*TreeNode, 0, len(nums))
	root := &TreeNode{}
	nodes = append(nodes, root)
	for i := 0; len(stack) > 0; {
		left := stack[0]
		right := stack[1]
		stack = append([]int{}, stack...)
		mid := left + (right - left) / 2
		node := nodes[i]
		node.Val = nums[mid]
		l, r := left, mid-1
		if l <= r {
			node.Left = &TreeNode{}
			nodes = append(nodes, node.Left)
			stack = append(stack, l, r)
		}
		l, r = mid+1, right
		if l <= r {
			node.Right = &TreeNode{}
			nodes = append(nodes, node.Right)
			stack = append(stack, l, r)
		}
		i++
	}
	return root
}
