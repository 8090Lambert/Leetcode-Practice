package __reverse

import (
	"math"
)

func reverse(x int) int {
	intMin, intMax := int(math.Pow(-2, 31))-1, int(math.Pow(2, 31)) - 1
	result := 0
	for x != 0 {
		single := x % 10
		if intMax - (result * 10) < single || intMin - (result * 10) > single {
			return 0
		}
		x /= 10
		result = result * 10 + single
	}
	return result
}
