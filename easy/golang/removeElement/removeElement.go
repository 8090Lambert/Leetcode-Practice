package removeElement

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	current := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[current] = nums[i]
			current+=1
		}
	}

	return current
}

func removeElementDuplicate(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	start, end := 0, len(nums)
	for start < end {
		if nums[start] == val {
			nums[start] = nums[end-1]
			end--
		} else {
			start++
		}
	}

	return start
}
