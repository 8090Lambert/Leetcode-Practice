package isSymmetric

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func IsSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	//return IsDouble(root.Left, root.Right);
	nodes := []*TreeNode{
		root.Left, root.Right,
	}
	for len(nodes) > 0 {
		lNode := nodes[0]
		rNode := nodes[1]
		nodes := nodes[2:]
		if lNode.Val != rNode.Val {
			return false
		}
		nodes = append(nodes, lNode.Left, rNode.Right, lNode.Right, rNode.Left)
	}

	return true
}

func IsDouble(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil && left.Val == right.Val{
		return IsDouble(left.Left, right.Right) && IsDouble(left.Right, right.Left)
	} 

	return false
}