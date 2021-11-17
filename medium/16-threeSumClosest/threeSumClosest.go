package _6_threeSumClosest

import (
	"math"
	"sort"
)

func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums)
	count := len(nums)
	if count < 3 {
		return 0
	} else if count == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < count; i++ {
		left := i + 1
		right := count - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
				res = sum
			}
			if sum == target {
				return target
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < 3; i++ {
		res += nums[i]
	}
	count := len(nums)
	for i := 0; i < count && count > 3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i+1
		right := count-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if int(math.Abs(float64(target-sum))) < int(math.Abs(float64(target-res))) {
				res = sum
			}
			if sum == target {
				return target
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	count := len(nums)
	res := 0
	for i := 0; i < 3; i++ {
		res += nums[i]
	}
	for i := 0; i < count; i++ {
		left := i+1
		right := count-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if int(math.Abs(float64(target-sum))) < int(math.Abs(float64(target-res))) {
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