package _89_rotate

func Rotate(nums []int, k int) {
	count := len(nums)
	if k%count == 0 || count == 0 || k == 0 {
		return
	}
	k %= count
	tmp := nums
	nums = append(nums[count-k:], nums[0:count-k]...)
	tmp = append(tmp[:0], nums...)
}

func rotate(nums []int, k int)  {
	k %= len(nums)
	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
