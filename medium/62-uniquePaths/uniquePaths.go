package _2_uniquePaths

func UniquePaths (m int, n int) int {
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

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[j-1][i]
		}
	}
	return dp[m-1][n-1]
}