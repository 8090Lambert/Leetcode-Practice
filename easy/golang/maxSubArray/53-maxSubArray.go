package maxSubArray

func maxSubArray(nums []int) int {
	current := 0
	sum := nums[0]
	for index := 0; index < len(nums); index++ {
		if current < 0 {
			current = nums[index]
		} else {
			current += nums[index]
		}
		if current > sum {
			sum = current
		}
	}

	return sum
}
