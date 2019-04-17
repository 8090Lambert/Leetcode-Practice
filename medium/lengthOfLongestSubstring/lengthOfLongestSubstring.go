package lengthOfLongestSubstring

import "math"

func LengthOfLongestSubstring(s string) int {
	count, res := len(s), 0
	m := make(map[string]int)
	for index, i := 0, 0; index < count; index++ {
		if _, ok := m[string(s[index])]; ok {
			i = int(math.Max(float64(m[string(s[index])]), float64(i)))
		}
		res = int(math.Max(float64(res), float64(index-i+1)))
		m[string(s[index])] = index + 1
	}

	return res
}

func LengthOfLongestSubstring1(s string) int {
	sArr := make(map[byte]byte)
	count := len(s)
	ret, start, end := 0, 0, 0
	for start < count && end < count {
		if _, ok := sArr[s[end]]; !ok {
			sArr[s[end]] = s[end]
			end++
			ret = int(math.Max(float64(ret), float64(end-start)))
		} else {
			delete(sArr, s[start])
			start++
		}
	}
	return ret
}
