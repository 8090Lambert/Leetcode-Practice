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
	l := len(nums) - 2
	for l >= 0 && nums[l] > nums[l+1] {
		l--
	}
	if l >= 0 {
		r := len(nums)-1
		for r >= 0 && nums[l] >= nums[r] {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	reverse1(&nums, l+1)
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
