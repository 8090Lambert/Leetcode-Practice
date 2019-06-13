package search

func search(nums []int, target int) int {
	return searchLoop(nums, 0, len(nums)-1, target)
}

func searchRecursion(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] < nums[right] {
		if nums[mid] < target && target <= nums[right] {
			return searchRecursion(nums, mid+1, right, target)
		} else {
			return searchRecursion(nums, left, mid-1, target)
		}
	} else {
		if nums[mid] > target && target >= nums[left] {
			return searchRecursion(nums, left, mid-1, target)
		} else {
			return searchRecursion(nums, mid+1, right, target)
		}
	}
}

func searchLoop(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[right] {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

func search81 (nums []int, target int) bool {
	count := len(nums)
	if count == 0 {
		return false
	}

	left := 0
	right := count - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[right] && nums[mid] == nums[left] {
			left++
			right--
		} else if nums[mid] <= nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return false
}
