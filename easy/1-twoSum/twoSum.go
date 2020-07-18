package __twoSum

func twoSum(nums []int, target int) []int {
	exists := make(map[int]int, len(nums))
	keys := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		if _, ok := exists[need]; ok {
			keys = append(keys, i, exists[need])
		}
		exists[nums[i]] = i
	}

	return keys
}


// 167 Question
func twoSum1(numbers []int, target int) []int {
	mp := make(map[int]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(numbers); i++ {
		need := target - numbers[i]
		if index, ok := mp[need]; ok {
			res = append(res, index, i+1)
			break
		} else {
			mp[need] = i+1
		}
	}
	return res
}