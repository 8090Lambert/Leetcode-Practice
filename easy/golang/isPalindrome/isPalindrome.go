package isPalindrome

import (
	"strconv"
	"strings"
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
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	begin, times, length := 0, 0, len(strconv.Itoa(x))
	if length%2 == 1 {
		times = length/2 + 1
	} else {
		times = length / 2
	}
	res := 0
	for begin < times {
		res = res*10 + x%10
		if begin+1 != times {
			x /= 10
		}
		begin++
	}

	return res == x || res == x/10
}

// 回文字符串
func IsPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}
	strArr := strings.Split(strings.ToLower(s), "")
	start := 0
	end := len(strArr) - 1
	for start < end {
		for isValid(strArr[start]) == false && start < end {
			start++
		}
		for isValid(strArr[end]) == false && start < end {
			end--
		}
		if strArr[start] == strArr[end] {
			start++
			end--
			continue
		} else {
			return false
		}
	}

	return true
}

func isValid(s string) bool {
	return (s >= "a" && s <= "z") || (s >= "0" && s <= "9")
}
