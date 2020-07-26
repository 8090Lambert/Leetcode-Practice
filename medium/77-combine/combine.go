package _7_combine

var res [][]int

func Combine(n int, k int) [][]int {
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


func combine(n int, k int) [][]int {
	res := [][]int{}
	back(1, k, n, []int{}, &res)
	return res
}

func back(start, k, n int, item []int, res *[][]int) {
	if len(item) == k {
		*res = append(*res, item)
		return
	}
	for i := start; i <= n; i++ {
		tmp := append([]int{}, item...)
		tmp = append(tmp, i)
		back(i+1, k, n, tmp, res)
	}
}