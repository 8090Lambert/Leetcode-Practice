package _5_sortColors

func SortColors(nums []int)  {
	current, second, third := 0, 0, len(nums)-1

	for current <= third {
		if nums[current] == 0 {
			nums[second],nums[current] = nums[current], nums[second]
			current++
			second++
		} else if nums[current] == 2 {
			nums[third],nums[current] = nums[current], nums[third]
			third--
		} else {
			current++
		}
	}
}

func sortColors(nums []int)  {
	current, second, third := 0, 0, len(nums)-1
	for current <= third {
		if nums[current] == 0 {
			nums[second], nums[current] = nums[current], nums[second]
			current++
			second++
		} else if nums[current] == 2 {
			nums[current], nums[third] = nums[current], nums[third]
			third--
		} else {
			current++
		}
	}
}
