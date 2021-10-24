package maxArea

import "math"

func MaxArea(height []int) int {
	start := 0
	end := len(height) - 1
	max := 0
	for start < end {
		current := int(math.Min(float64(height[start]), float64(height[end]))) * (end - start)
		max = int(math.Max(float64(current), float64(max)))
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return max
}




func maxArea (height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		minHeight := int(math.Min(float64(height[left]), float64(height[right])))
		cur := minHeight * (right-left)
		if cur > res {
			res = cur
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return res
}