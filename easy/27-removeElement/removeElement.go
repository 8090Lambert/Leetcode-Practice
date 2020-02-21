package _7_removeElement

func removeElement(nums []int, val int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}
	
	prev := 0
	for i := 0; i < count; i++ {
		if nums[i] != val {
			nums[prev] = nums[i]
			prev++
		}
	}
	
	return prev
}

// two pointer
func removeElementMethod2(nums []int, val int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}
	
	start, end := 0, len(nums)-1
	for start <= end {
		if nums[start] == val {
			nums[start] = nums[end]
			end -=1
		} else {
			start++
		}
	}
	return start
}
