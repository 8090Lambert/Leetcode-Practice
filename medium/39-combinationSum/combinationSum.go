package _9_combinationSum

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

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	collect := make([][]int, 0)
	back(candidates, 0, target, []int{}, &collect)
	return collect
}

func back(candidates []int, start, target int, item []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, item)
		return
	}
	for i := start; i < len(candidates) && target >= candidates[i]; i++ {
		temp := make([]int, 0, len(item)+1)
		temp = append(temp, item...)
		temp = append(temp, candidates[i])
		back(candidates, i, target - candidates[i], temp, res)
	}
}
