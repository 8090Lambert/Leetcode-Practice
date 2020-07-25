package _7_permuteUnique

func PermuteUnique(nums []int) [][]int {
	res := [][]int{}
	backTrack(nums, &res, 0)

	return res
}

func backTrack(nums []int, res *[][]int, start int) {
	if len(nums) == start {
		tmp := append([]int{}, nums...)
		*res = append(*res, tmp)
		return
	}

	unique := map[int]bool{}
	for i := start; i < len(nums); i++ {
		if _, ok := unique[nums[i]]; ok {
			continue
		}
		nums[i], nums[start] = nums[start], nums[i]
		backTrack(nums, res, start+1)
		nums[start], nums[i] = nums[i], nums[start]
		unique[nums[i]] = true
	}
}
