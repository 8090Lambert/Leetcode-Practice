package _2_generateParenthesis

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

func generateParenthesis(n int) []string {
	res := new([]string)
	backDel(res, "", n, n)
	return *res
}

func backDel(res *[]string, tmp string, start, close int) {
	if start == 0 && close == 0 {
		*res = append(*res, tmp)
		return
	}
	if start > 0 {
		backDel(res, tmp+"(", start-1, close)
	}
	if close > start {
		backDel(res, tmp+")", start, close-1)
	}
}