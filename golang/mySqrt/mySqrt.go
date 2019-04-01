package mySqrt

import "fmt"

func MySqrt(x int) int {
	result := x
	for result > x / result {
		result = (result +  x / result) / 2
	}

	fmt.Println(result)
	return result
}