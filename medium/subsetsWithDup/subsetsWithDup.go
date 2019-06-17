package subsetsWithDup

import (
	"sort"
	"strconv"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	backTrack(0, nums, []int{}, &res)

	return res
}


func backTrack (start int, nums []int, item []int, res *[][]int) {
	*res = append(*res, item)

	pre := "u"
	for i := start; i < len(nums); i++ {
		if pre == strconv.Itoa(nums[i]) {
			continue
		}
		pre = strconv.Itoa(nums[i])
		item = append(item, nums[i])
		backTrack(i+1, nums, item, res)
		item = append([]int{}, item[:len(item)-1]...)
	}
}