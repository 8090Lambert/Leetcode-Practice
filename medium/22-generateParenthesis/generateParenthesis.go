package generateParenthesis

func GenerateParenthesis(n int) []string {
	strArr := make([]string, 0)
	back(&strArr, "", 0, 0, n)
	return strArr
}

func back(strArr *[]string, cur string, open, close, max int) {
	if len(cur) != max*2 {
		if open < max {
			back(strArr, cur+"(", open+1, close, max)
		}

		if close < open {
			back(strArr, cur+")", open, close+1, max)
		}
	} else {
		*strArr = append(*strArr, cur)
	}
}
