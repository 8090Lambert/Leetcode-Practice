package subsets

func subsets(nums []int) [][]int {
	res := [][]int{}
	backTrack(0, nums, []int{}, &res)
	return res
}

func backTrack (offset int, nums []int, item []int, res *[][]int) {
	*res = append(*res, item)

	for i := offset; i < len(nums); i ++ {
		item = append(item, nums[i])
		backTrack(i+1, nums, item, res)
		item = append([]int{}, item[:len(item)-1]...)
	}
}
