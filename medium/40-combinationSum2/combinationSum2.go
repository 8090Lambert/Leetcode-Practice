package _0_combinationSum2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	backTrack(candidates, []int{}, target, 0, &res)
	return res
}


func backTrack(candidates, item []int, target, offset int, res *[][]int) {
	if target == 0 {
		*res = append(*res, item)
		return
	}
	for i := offset; i < len(candidates) && target >= candidates[i]; i++ {
		origin := i
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
		tmp := make([]int, 0, len(item)+1)
		tmp = append(tmp, item...)
		tmp = append(tmp, candidates[i])
		backTrack(candidates, tmp, target-candidates[i], origin+1, res)
	}
}
