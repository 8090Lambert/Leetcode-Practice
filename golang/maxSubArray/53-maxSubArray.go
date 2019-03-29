package maxSubArray

func Main(nums []int) int {
	current := 0
	res := nums[0]

	for index := 0; index < len(nums); index++ {
		if current > 0 {
			current += nums[index]
		} else {
			current = nums[index]
		}

		if current > res {
			res = current
		}
	}
	return res
}

func MainCopy(nums []int) int {
	current := nums[0]
	sum := nums[0]
	for index := 1; index < len(nums); index++ {
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
