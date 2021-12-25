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


func buildTreePre(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTreePre(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTreePre(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}


func buildTreePost(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if root.Val == inorder[i] {
			break
		}
	}
	root.Left = buildTreePost(inorder[:i], postorder[:i])
	root.Right = buildTreePost(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}