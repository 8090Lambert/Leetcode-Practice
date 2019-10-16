package reverse

import "math"

func reverse(x int) int {
	intMin, intMax := int(math.Pow(-2, 31)) - 1, int(math.Pow(2, 31))-1
	res := 0
	for x != 0 {
		tail := x % 10
		x /= 10
		if res > (intMax - tail) / 10  || res < (intMin - tail) / 10 {
			return 0
		}
		res = res*10 + tail
	}

	return res
}
