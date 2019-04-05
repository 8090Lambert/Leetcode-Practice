package isPalindrome

import (
	"fmt"
	"strconv"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	origin := x
	var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	return result == origin
}

func IsPalindrome1(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	begin := 0
	length := len(strconv.Itoa(x))
	var times int
	if length % 2 == 1 {
		times = length / 2 + 1
	} else {
		times = length / 2
	}

	res := 0
	//fmt.Println(begin, times)
	for begin < int(times) {
		res = res*10 + x%10
		if begin + 1 != times {
			x /= 10
		}
		begin++
	}

	fmt.Println(res, x)
	fmt.Println(res == x || res == x / 10)
	return res == x || res == x / 10
}
