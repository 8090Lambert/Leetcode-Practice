package _5_searchInsert

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}



func searchInsert1(nums []int, target int) int {
	count := len(nums)
	if count <= 0 {
		return 0
	}
	start, end := 0, count-1
	for start <= end {
		mid := start + (end - start) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}