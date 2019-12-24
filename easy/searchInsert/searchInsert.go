package searchInsert

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)-1
	if nums[end] < target {
		return end+1
	}

	for start <= end {
		mid := start + (end-start) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return start
}
