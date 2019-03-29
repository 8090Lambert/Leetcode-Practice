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
