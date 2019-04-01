package mySqrt

func MySqrt(x int) int {
	result := x
	for result > x / result {
		result = (result +  x / result) / 2
	}

	return result
}