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


func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return  ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	isCommonPrefix := func (length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 0; i < count; i++ {
			if str0 != strs[i][:length] {
				return false
			}
		}
		return true
	}
	min := len(strs[0])
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}
	start, end := 0, min
	for start < end {
		mid := (end - start + 1) / 2 + start
		if isCommonPrefix(mid) {
			start = mid
		} else {
			end = mid - 1
		}
	}
	return strs[0][:start]
}



func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		for len(prefix) > 0 && !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return prefix
			}
		}
	}
	return prefix
}