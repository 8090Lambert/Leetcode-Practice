package searchRange

import (
	"math"
)

func SearchRange(nums []int, target int) []int {
	position, start, end := -1, 0, len(nums)
	left, right := 0, 0
	for start <= end {
		half := int(math.Floor(float64((start + end) / 2)))
		if half < 0 || half >= len(nums) {
			break
		}
		if nums[half] == target {
			position = half
			left, right = half, half
			for true {
				if left-1 >= 0 && nums[left-1] == target {
					left = left - 1
				} else if right+1 < len(nums) && nums[right+1] == target {
					right = right + 1
				} else {
					break
				}
			}
			break
		} else if nums[half] < target {
			start = half + 1
		} else {
			end = half - 1
		}
	}

	if position == -1 {
		return []int{-1, -1}
	}

	return []int{left, right}
}
