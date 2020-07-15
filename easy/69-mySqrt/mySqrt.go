package _9_mySqrt

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	result := x
	for result > x /result {
		result = (result + x/result) / 2
	}

	return result
}

func mySqrtTwoSplit(x int) int {
	if x == 0 {
		return 0
	}
	left := 0
	right := x
	for left < right {
		mid := (left + right+1) >> 1
		if mid * mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}