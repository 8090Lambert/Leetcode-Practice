package _12_hasPathSum

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


func hasPathSumLoop(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	stack := []*TreeNode{root}
	iset := []int{targetSum}
	for len(stack) != 0 {
		count := len(stack)
		for i := 0; i < count; i++ {
			num := iset[i]
			cur := stack[i]
			if cur.Val == num && cur.Left == nil && cur.Right == nil {
				return true
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
				iset = append(iset, num - cur.Val)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				iset = append(iset, num - cur.Val)
			}
		}
		stack = append([]*TreeNode{}, stack[count:]...)
		iset = append([]int{}, iset[count:]...)
	}

	return false
}