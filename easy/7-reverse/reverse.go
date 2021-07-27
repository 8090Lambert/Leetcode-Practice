package __reverse

func reverse(x int) int {
	min, max := -1 << 31, (1 << 31) - 1
	res := 0
	for x != 0 {
		end := x % 10
		if max - end < res * 10 || min - end > res * 10 {
			return 0
		}
		res = res * 10 + end
		x /= 10
	}
	return res
}