package _15_findKthLargest

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	count := len(nums)
	target := count-k
	left, right := 0, count-1
	for {
		pos := partition(nums, left, right)
		if pos == target {
			return nums[pos]
		} else if pos < target {
			left = pos+1
		} else {
			right = pos-1
		}
	}
}

func partition(nums []int, left, right int) int {
	if left < right {
		randPos := rand.Int() % (right-left) + left + 1
		nums[left], nums[randPos] = nums[randPos], nums[left]
	}
	x := nums[left]
	j := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < x {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[j], nums[left] = nums[left], nums[j]
	return j
}
