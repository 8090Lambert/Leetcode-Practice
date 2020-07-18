package _69_majorityElement

func MajorityElement(nums []int) int {
	counter, ret := 0, 0
	count := len(nums)
	for index := 0; index < count; index++ {
		if counter == 0 {
			ret = nums[index]
			counter = 1
		} else if ret == nums[index] {
			counter++
		} else {
			counter--
		}
	}
	return ret
}

func majorityElement(nums []int) int {
	sumMap := make(map[int]int, 0)
	count := len(nums)
	for i := 0; i < count; i++ {
		if sum, ok := sumMap[nums[i]]; ok {
			if sum+1 >= (count >> 1) {
				return nums[i]
			}
			sumMap[nums[i]] += 1
		} else {
			sumMap[nums[i]] = 1
		}
	}
	return 0
}
