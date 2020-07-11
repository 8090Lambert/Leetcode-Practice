package lastword

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	if !strings.Contains(s, " ") {
		return len(s)
	}

	s = strings.Trim(s, " ")
	offset := 1
	strSlinces := strings.Split(s, " ")
	count := len(strSlinces)
	for strSlinces[len(strSlinces) - offset] == "" && offset <= count {
		offset++
	}
	return len(strSlinces[len(strSlinces) - offset])
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == ' ' && count == 0 {
			continue
		} else if s[i] == ' ' && count != 0 {
			break
		} else {
			count++
		}
	}

	return count
}











func LengthOfLastWordStep(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	count := 0
	for i := length - 1; i > -1; i-- {
		if s[i] == ' ' && count == 0 {
			continue
		} else if s[i] == ' ' && count != 0 {
			break
		} else {
			count++
		}
	}
	return count
}