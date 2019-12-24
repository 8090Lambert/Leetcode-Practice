package hasPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum-root.Val == 0
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}


func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	stack := []*TreeNode{root}
	total := []int{sum}
	for len(stack) > 0 {
		count := len(stack)
		current := stack[count-1]
		stack = append([]*TreeNode{}, stack[:count-1]...)
		target := total[len(total)-1]
		total = append([]int{}, total[:len(total)-1]...)
		if target == current.Val && current.Left == nil && current.Right == nil {
			return true
		}
		if current.Left != nil {
			total = append(total, target - current.Val)
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			total = append(total, target - current.Val)
			stack = append(stack, current.Right)
		}
	}
	return false
}