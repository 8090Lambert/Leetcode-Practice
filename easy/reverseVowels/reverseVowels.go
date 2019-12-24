package reverseVowels

func reverseVowels (s string) string {
	all := map[string]struct{}{
		"a": struct{}{}, "e": struct{}{}, "i": struct{}{}, "o": struct{}{}, "u": struct{}{},
		"A": struct{}{}, "E": struct{}{}, "I": struct{}{}, "O": struct{}{}, "U": struct{}{},
	}
	left, right := 0, len(s)-1
	sBytes := []byte(s)
	for left < right {
		_, okLeft := all[string(s[left])]
		_, okRight := all[string(s[right])]
		if okLeft && okRight {
			sBytes[left], sBytes[right] = s[right], s[left]
			left++
			right--
		}
		if !okLeft {
			left++
		}
		if !okRight {
			right--
		}
	}

	return string(sBytes)
}
