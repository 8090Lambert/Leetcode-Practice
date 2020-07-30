package _0_subsetsWithDup

import (
	"sort"
	"strconv"
)

func SubsetsWithDup(nums []int) [][]int {
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

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	back(nums, 0, []int{}, &res)
	return res
}

func back(nums []int, start int, item []int, res *[][]int) {
	*res = append(*res, item)
	var uniq string
	for i := start; i < len(nums); i++ {
		if uniq == strconv.Itoa(nums[i]) {
			continue
		}
		uniq = strconv.Itoa(nums[i])
		tmp := make([]int, 0, len(item)+1)
		tmp = append(tmp, item...)
		tmp = append(item, nums[i])
		back(nums, i+1, tmp, res)
	}
}