package removeDuplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev := 0
	for i := 1; i < len(nums); i++ {
		if nums[prev] != nums[i] {
			prev++
			nums[prev] = nums[i]
		}
	}
	return prev + 1
}
























func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev := 0
	for i := 1; i < len(nums); i++ {
		if nums[prev] != nums[i] {
			prev++
			nums[prev] = nums[i]
		}
	}

	return prev + 1
}