package _05_buildTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree (preorder []int, inorder []int) *TreeNode {
	preLen := len(preorder)
	inLen := len(inorder)
	return backTrackPre(preorder, 0, preLen-1, inorder, 0, inLen-1)
}

/**
	前序遍历 + 中序遍历
 */
func backTrackPre (preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	
	rootVal := preorder[preStart]
	l := inStart
	for l <= inEnd && inorder[l] != rootVal {
		l++
	}
	
	root := &TreeNode{Val:rootVal}
	root.Left = backTrackPre(preorder, preStart+1, preStart+l-inStart, inorder, inStart, l-1)
	root.Right = backTrackPre(preorder, preStart+l-inStart+1, preEnd, inorder, l+1, inEnd)
	
	return root
}





/**
	后序遍历 + 中序遍历
 */
func backTrackPost (inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd || postStart > postEnd {
		return nil
	}
	
	rootVal := postorder[postEnd]
	l := inStart
	for l <= inEnd && inorder[l] != rootVal {
		l++
	}
	
	root := &TreeNode{Val:rootVal}
	root.Left = backTrackPost(inorder, inStart, l-1, postorder, postStart, postStart + l - inStart - 1)
	root.Right = backTrackPost(inorder, l+1, inEnd, postorder, postStart+l-inStart, postEnd-1)
	
	return root
}