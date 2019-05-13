package search

func search(nums []int, target int) int {
	return searchM(nums, 0, len(nums)-1, target)
}

func searchM(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] < nums[right] {
		if nums[mid] < target && target <= nums[right] {
			return searchM(nums, mid+1, right, target)
		} else {
			return searchM(nums, left, mid-1, target)
		}
	} else {
		if nums[mid] > target && target >= nums[left] {
			return searchM(nums, left, mid-1, target)
		} else {
			return searchM(nums, mid+1, right, target)
		}
	}
}
