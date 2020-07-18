package _70_trailingZeroes

func TrailingZeroes(n int) int {
	count := 0
	for n > 1 {
		count += n / 5
		n /=5 
	}

	return count
}


// 仅看5的个数就可以了
func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		count += n / 5
		n /= 5
	}
	return count
}
