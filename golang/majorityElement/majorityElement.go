package majorityElement

import "fmt"

func MajorityElement(nums []int) int {
	count := 0
	ret := nums[0]
	length := len(nums)
	for index := 1; index < length; index++ {
		if ret == nums[index] {
			count++
		} else {
			count--
		}
	}

	return ret
}