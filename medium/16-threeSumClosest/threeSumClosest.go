package _6_threeSumClosest

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	count := len(nums)
	if count < 3 {
		return 0
	}
	if count == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	close := nums[0] + nums[1] + nums[2]
	for index := 0; index < count-1; index++ {
		left := index + 1
		right := count - 1
		for left < right {
			sum := nums[index] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-close)) {
				close = sum
			}
			if sum == target {
				return target
			} else if sum > target{
				right--
			} else {
				left++
			}
		}
	}

	return close
}



func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums)
	count := len(nums)
	if count < 3 {
		return 0
	} else if count == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < count-1; i++ {
		left := i + 1
		right := count-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
				res = sum
			}
			if sum == target {
				return target
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
