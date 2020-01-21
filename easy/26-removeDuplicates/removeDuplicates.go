package _6_removeDuplicates

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	current, count := nums[0], 1
	for i := 2; i < len(nums); i++ {
		if nums[i] != current {
			current = nums[i]
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

// 0 1 1 1 2 2 3 -> 0 1 2 3
