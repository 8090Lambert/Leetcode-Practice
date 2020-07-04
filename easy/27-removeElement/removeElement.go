package _7_removeElement

func removeElement(nums []int, val int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}
	start, end := 0, count-1
	for start <= end {
		if nums[start] == val {
			nums[start] = nums[end]
			end--
		} else {
			start++
		}
	}
	return start
}
