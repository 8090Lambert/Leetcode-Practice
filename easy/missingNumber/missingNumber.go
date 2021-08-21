package missingNumber

func missingNumber(nums []int) int {
	total := len(nums)
	for i := 0; i < len(nums); i++ {
		tmp := i ^ nums[i]
		total ^= tmp
	}
	
	return total
}

func missingNumberMath(nums []int) int {
	total := (0 + len(nums)) * (len(nums)) / 2
	for i := 0; i < len(nums); i++ {
		total -= nums[i]
	}
	return total
}
