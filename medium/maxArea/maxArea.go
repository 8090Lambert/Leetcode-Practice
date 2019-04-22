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
