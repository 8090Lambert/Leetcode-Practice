package isHappy

import "math"

func isHappy (n int) bool {
	if n == 1 || n == 7 {
		return true
	}
	
	if n < 10 {
		return false
	}
	
	sum := 0
	for n >= 1 {
		tmp := n % 10
		sum += int(math.Pow(float64(tmp), float64(2.0)))
		n /= 10
	}
	
	return isHappy(sum)
}
