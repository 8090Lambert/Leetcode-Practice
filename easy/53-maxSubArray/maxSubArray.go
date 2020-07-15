package _3_maxSubArray

func maxSubArray(nums []int) int {
	cur := 0
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if cur < 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}
		if cur > res {
			res = cur
		}
	}
	return res
}


func maxSubArrayOther(nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < count; i++ {
		if nums[i-1] + nums[i] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
