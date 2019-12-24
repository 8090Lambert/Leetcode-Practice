package trailingZeroes

func TrailingZeroes(n int) int {
	count := 0
	for n > 1 {
		count += n / 5
		n /=5 
	}

	return count
}
