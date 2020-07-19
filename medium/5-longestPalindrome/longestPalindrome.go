package __longestPalindrome

import (
	"math"
)

func LongestPalindrome(s string) string {
	count := len(s)
	if s == "" || count <= 1{
		return s
	}
	start, end := 0, 0
	for index := 0; index < count; index++ {
		len1 := expandAroundCenter(s, index, index)
		len2 := expandAroundCenter(s, index, index+1)
		len := int(math.Max(float64(len1), float64(len2)))
		if len > end - start {
			start = index - (len - 1) / 2
			end = index + len / 2
		}
	}
	return s[start: end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
    for L >= 0 && R < len(s) && s[L] == s[R] {
        L = L - 1
        R = R + 1
	}
	//fmt.Println(R, L)
    return R - L - 1
}


func longestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := extend(s, i, i)
		l2, r2 := extend(s, i, i+1)
		if r1 - l1 > end - start {
			start, end = l1, r1
		}
		if r2 - l2 > end - start {
			start, end = l2, r2
		}
	}
	return s[start:end+1]
}

func extend(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l = l-1
		r = r+1
	}
	return l + 1, r - 1
}