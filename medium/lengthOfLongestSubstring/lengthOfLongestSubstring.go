package lengthOfLongestSubstring

import (
	"math"
)

/**
	滑动窗口优化，每次找重复字符串的索引位置
 */
func LengthOfLongestSubstring(s string) int {
	res, count := 0, len(s)
	window := map[byte]int{}
	j := 0
	for i := 0; i < count; i++ {
		if _, ok := window[s[i]]; ok{
			j = int(math.Max(float64(window[s[i]]), float64(j)))
		}
		res = int(math.Max(float64(res), float64(i-j+1)))
		window[s[i]] = i+1
	}

	return res
}

/**
	滑动窗口，每次比较end对应的字符串 start - end 之间
 */
func lengthOfLongestSubstringWithWindow(s string) int {
	res, count := 0, len(s)
	windows := map[byte]struct{}{}
	start, end := 0, 0
	for start < count && end < count {
		if _, ok := windows[s[end]]; !ok {
			windows[s[end]] = struct{}{}
			end++
			res = int(math.Max(float64(res), float64(end - start)))
		} else {
			delete(windows, s[start])
			start++
		}
	}
	return res
}
