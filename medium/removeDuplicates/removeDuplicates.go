package removeDuplicates

func RemoveDuplicates (nums []int) int {
	count := len(nums)
	if count <= 1 {
		return count
	}

	current := 1
	for i := 2; i < count; i++ {
		if nums[i] != nums[current-1] {
			current++
			nums[current] = nums[i]
		}
	}

	return current+1
}