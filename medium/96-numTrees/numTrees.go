package _6_numTrees

/**
	计算公式：G(n) = G(i-1) * G(n-i), n为个数
 */
func numTrees(n int) int {
	count := make([]int, n+1)
	count[0], count[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			count[i] += count[j-1] * count[i-j]
		}
	}
	return count[n]
}


func numTreesOther(n int) int {
	count := make([]int, n+1)
	count[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			count[i] += count[j-1] * count[i-j]
		}
	}
	return count[n]
}