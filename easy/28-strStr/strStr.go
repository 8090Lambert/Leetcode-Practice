package _8_strStr

func strStr(haystack string, needle string) int {
	hlength, nlength := len(haystack), len(needle)
	if haystack == needle || nlength == 0 {
		return 0
	}
	if hlength < nlength {
		return - 1
	}
	for i := 0; i < hlength-nlength; i++ {
		if haystack[i:i+nlength] == needle {
			return i
		}
	}

	return -1
}



func StrStr(haystack string, needle string) int {
	sLength, nLength := len(haystack), len(needle)
	if haystack == needle || nLength == 0 {
		return 0
	}
	if sLength < nLength  {
		return -1
	}
	next := getNext(needle)
	i, j := 0, 0
	for i < sLength && j < nLength {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == nLength {
		return i - j
	} else {
		return -1
	}
}


func getNext (needle string) []int {
	length := len(needle)
	next := make([]int, length, length)
	next[0] = -1
	j, k := 0, -1
	for j < length - 1 {
		if k == -1 || needle[j] == needle[k] {
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}
	}

	return next
}
