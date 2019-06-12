package combine

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	backTrack(1, k, n, []int{})
	return res
}

func backTrack (first, k, n int, set []int) {
	if len(set) == k {
		res = append(res, set)
		return
	}

	for i := first; i <= n; i++ {
		tmp := append([]int{}, set...)
		tmp = append(tmp, i)
		backTrack(i+1,k,n,tmp)
	}
}
