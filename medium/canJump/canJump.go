package canJump

func CanJump(nums []int) bool {
	lastPosition := len(nums) - 1
	for i := lastPosition; i >= 0; i-- {
		if i + nums[i] >= lastPosition {
			lastPosition = i
		}
	}

	return lastPosition == 0
}
