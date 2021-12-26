package _15_findKthLargest

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	count := len(nums)
	left, right := 0, count-1
	target := count-k
	for {
		pos := partition(nums, left, right)
		if pos == target {
			return nums[target]
		} else if pos > target {
			right=pos-1
		} else {
			left=pos+1
		}
	}
}

// x=0, j=0, [3,1,5,6,2,4]
// i=1, j=1, [3,1,5,6,2,4]
// i=2, j=1, [3,1,5,6,2,4]
// i=3, j=1, ...
// i=4, j=2, [3,1,2,6,5,4]
// i=5, j=2
func partition(nums []int, left, right int) int {
	if left < right {
		randPos := rand.Int() % (right-left) + left + 1
		nums[randPos], nums[left] = nums[left], nums[randPos]
	}
	x := nums[left]
	j := left
	for i := left+1; i <= right; i++ {
		if nums[i] < x {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[j], nums[left] = nums[left], nums[j]
	return j
}