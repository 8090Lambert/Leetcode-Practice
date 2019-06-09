package uniquePaths

func uniquePaths (m int, n int) int {
	res := []int{}
	for index := 0; index < n; index++ {
		res = append(res, 1)
	}
	
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[j] += res[j-1]
		}
	}
	
	return res[n-1]
}