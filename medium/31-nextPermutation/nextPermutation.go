package _1_nextPermutation

func NextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		(nums)[i], (nums)[j] = (nums)[j], (nums)[i]
	}

	reverse(&nums, i+1)
}

func reverse(nums *[]int, start int) {
	i := start
	j := len(*nums) - 1
	for i < j {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		i++
		j--
	}
}





func nextPermutation(nums []int)  {
	left := len(nums)-2
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left >= 0 {
		right := len(nums)-1
		for right >= 0 && nums[left] >= nums[right] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	reverse1(&nums, left+1)
}

func reverse1(nums *[]int, start int) {
	l := start
	r := len(*nums)-1
	for l < r {
		(*nums)[l], (*nums)[r] = (*nums)[r], (*nums)[l]
		l++
		r--
	}
}
