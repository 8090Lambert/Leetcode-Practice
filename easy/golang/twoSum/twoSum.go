package twoSum

func twoSum(nums []int, target int) []int {
	hasMap := make(map[int]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		if _, ok := hasMap[need]; !ok || len(hasMap) <= 0{
			hasMap[nums[i]] = i
			continue
		}
		res = append(res, hasMap[nums[i]], i)
	}

	return res
}