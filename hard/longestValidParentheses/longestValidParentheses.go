package longestValidParentheses

import "math"

func LongestValidParentheses(s string) int {
	stack := []int{-1}
	count := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			stack = append(stack, i)
		} else {
			stack = append([]int{}, stack[:len(stack) - 1]...)
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				count = int(math.Max(float64(count), float64(i-stack[len(stack)])))
			}
		}
	}
	return count
}