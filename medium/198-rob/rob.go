package _98_rob

import "math"

func rob(nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}
	if count == 1 {
		return nums[0]
	}
	first, second := nums[0], int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < count; i++ {
		third := int(math.Max(float64(first + nums[i]), float64(second)))
		first, second = second, third
	}
	return second
}
