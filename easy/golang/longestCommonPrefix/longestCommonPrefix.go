package longestCommonPrefix

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	count := len(strs)
	if count == 0 {
		return ""
	}
	if count == 1 {
		return strs[0]
	}
	commonPre := strs[0]
	for i := 1; i < count; i++ {
		for strings.Index(strs[i], commonPre) != 0 {
			commonPre = commonPre[:len(commonPre)-1]
			if commonPre == "" {
				return ""
			}
		}
	}
	return commonPre
}
