package _4_longestCommonPrefix

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func longestCommonPrefixOther(strs []string) string {
	count := len(strs)
	if count == 0 {
		return ""
	}
	if count == 1 {
		return strs[0]
	}
	prefix := strs[1]
	var flag bool
	for len(prefix) != 0 {
		flag = true
		for i := 1; i < count; i++ {
			if !strings.HasPrefix(strs[i], prefix) {
				flag = false
				prefix = prefix[:len(prefix)-1]
				break
			}
		}
		if flag {
			break
		}
	}
	return prefix
}