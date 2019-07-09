package golang

func isPerfectSquare (num int) bool {
	if num <= 1 {
		return true
	}

	left := 0
	right := num/2
	for left < right {
		mid := (left + right) / 2
		if mid >= num / mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left * left == num
}
