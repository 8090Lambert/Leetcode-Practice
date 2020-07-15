package removeDuplicates

func removeDuplicates(nums []int) int {
	count := len(nums)
	if count < 2 {
		return count
	}
	prev := 0
	for i := 1; i < count; i++ {
		if nums[prev] != nums[i] {
			prev++
			nums[prev] = nums[i]
		}
	}
	return prev+1
}