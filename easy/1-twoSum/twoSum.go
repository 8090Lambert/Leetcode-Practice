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