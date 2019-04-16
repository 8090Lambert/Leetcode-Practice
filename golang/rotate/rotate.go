package rotate

func Rotate(nums []int, k int) {
	count := len(nums)
	if k%count == 0 || count == 0 || k == 0 {
		return
	}
	k %= count
	tmp := nums
	nums = append(nums[count-k:], nums[0:count-k]...)
	for index := 0; index < count; index++ {
		tmp[index] = nums[index]
	}
}
