package _7_insert

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		origin := res[len(res)-1]
		target := intervals[i]
		if target[0] > origin[1] {
			res = append(res, target)
			continue
		}
		if target[1] > origin[1] {
			res[len(res)-1][1] = target[1]
		}
	}
	return res
}