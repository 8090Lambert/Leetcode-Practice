package merge

import (
	"sort"
)

func Merge (intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	count := len(intervals)
	for i := 1; i < count; i++ {
		if intervals[i][0] <= res[len(res)-1][0] {
			if intervals[i][1] <= res[len(res)-1][1] {
				if intervals[i][1] > res[len(res)-1][0] {
					res[len(res)-1][0] = intervals[i][0]
				} else {
					tmp := [][]int{intervals[i]}
					res = append(tmp, res...)
				}
			} else if intervals[i][1] > res[len(res)-1][1] {
				res[len(res)-1][0] = intervals[i][0]
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			if intervals[i][0] > res[len(res)-1][1] {
				res = append(res, intervals[i])
			} else {
				if intervals[i][1] > res[len(res)-1][1] {
					res[len(res)-1][1] = intervals[i][1]
				}
			}
		}
	}
	
	return res
}
