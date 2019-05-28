package permute

var res [][]int

func Permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	var backTrack func(left, ans []int)
	backTrack = func(left, ans []int) {
		if len(left) == 0 {
			res = append(res, ans)
			return
		}

		for key, value := range left {
			cleft := append([]int{}, left...)
			cans := append([]int{}, ans...)

			backTrack(append(cleft[:key], cleft[key+1:]...), append(cans, value))
		}
	}
	backTrack(nums, []int{})

	return res
}
