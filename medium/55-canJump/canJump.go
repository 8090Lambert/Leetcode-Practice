package _5_canJump

func CanJump(nums []int) bool {
	lastPosition := len(nums) - 1
	for i := lastPosition; i >= 0; i-- {
		if i + nums[i] >= lastPosition {
			lastPosition = i
		}
	}

	return lastPosition == 0
}

func canJump(nums []int) bool {
	lastP := len(nums) - 1
	for i := lastP; i >=0; i-- {
		if i + nums[i] > lastP {
			lastP = i
		}
	}
	return lastP == 0
}
