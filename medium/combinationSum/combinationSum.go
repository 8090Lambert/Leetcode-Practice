package combinationSum

import (
	"sort"
)

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	collect := backtrack(candidates, target, 0, []int{})
	return collect

}

func backtrack(candidates []int, target, i int, tmp []int) [][]int {
	if target == 0 {
		return [][]int{tmp}
	}

	res := [][]int{}
	for index := i; index < len(candidates) && target >= candidates[index]; index++ {
		temp := []int{}
		temp = append(temp, tmp...)
		temp = append(temp, candidates[index])
		allTmp := backtrack(candidates, target-candidates[index], index, temp)
		if allTmp != nil {
			res = append(res, allTmp...)
		}
	}

	return res
}
