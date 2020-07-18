package _36_singleNumber

func SingleNumber(nums []int) int {
	a := 1
	for index := 0; index < len(nums); index++ {
		a = a ^ nums[index]
	}

	return a ^ a
}
