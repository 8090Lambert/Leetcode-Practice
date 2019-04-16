package majorityElement

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

func MajorityElement1(nums []int) int {
	countMap := map[int]int{}
	count := len(nums)
	for index := 0; index < count; index++ {
		if _, ok := countMap[nums[index]]; ok {
			countMap[nums[index]] = countMap[nums[index]] + 1
		} else {
			countMap[nums[index]] = 1
		}
	}

	ret := 0
	for val, key := range countMap {
		if key >= count/2 {
			ret = val
		}
	}

	return ret
}
