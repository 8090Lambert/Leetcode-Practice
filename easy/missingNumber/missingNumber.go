package missingNumber

func missingNumber(nums []int) int {
	total := len(nums)
	for i := 0; i < len(nums); i++ {
		tmp := i ^ nums[i]
		total ^= tmp
	}
	
	return total
}
