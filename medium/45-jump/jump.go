package _5_jump

import "math"

func jump(nums []int) int {
	pos := len(nums)-1
	step := 0
	for pos > 0 {
		for i := 0; i < pos; i++ {
			if i + nums[i] >= pos {
				pos = i
				step++
				break
			}
		}
	}
	return step
}


func Jump(nums []int) int {
	count := len(nums)
	maxPos := 0
	end := 0
	step := 0
	for i := 0; i < count-1; i++ {
		maxPos = int(math.Max(float64(maxPos), float64(i+nums[i])))
		if i == end {
			end = maxPos
			step++
		}
	}
	return step
}
