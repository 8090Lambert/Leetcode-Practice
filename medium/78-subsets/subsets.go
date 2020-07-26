package _8_subsets

func Subsets(nums []int) [][]int {
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


func subsets(nums []int) [][]int {
	res := [][]int{}
	backTrace(0, nums, []int{}, &res)
	return res
}


func backTrace(start int, nums, item []int, res *[][]int) {
	*res = append(*res, item)
	for i := start; i < len(nums); i++ {
		tmp := make([]int, 0, len(item)+1)
		tmp = append(tmp, item...)
		tmp = append(tmp, nums[i])
		backTrace(i+1, nums, item, res)
	}
}