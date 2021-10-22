package _09_sortedListToBST

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
 	Val int
 	Left *TreeNode
	Right *TreeNode
}

func sortedListToBSTWithArr(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return buildTree(res)
}

func buildTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	start, end := 0, len(nums)-1
	mid := start + (end-start) >> 1
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = buildTree(nums[:mid])
	root.Right = buildTree(nums[mid+1:])
	return root
}



func sortedListToBST(head *ListNode) *TreeNode {
	return buildWithList(head, nil)
}

func buildWithList(head, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{slow.Val, nil, nil}
	root.Left = buildWithList(head, slow)
	root.Right = buildWithList(slow.Next, tail)
	return root
}