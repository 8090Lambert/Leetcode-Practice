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
	res := make([]string, 0)
	back2(&res, "", n, n)
	return res
}

func back1(res *[]string, cur string, o, c int) {
	if c == 0 && o == 0{
		*res = append(*res, cur)
		return
	}
	if o > 0 {
		back1(res, cur + "(", o-1, c)
	}
	if c > o {
		back1(res, cur + ")", o, c-1)
	}
}


func back2(res *[]string, cur string, o, c int) {
	if o == 0 && c == 0 {
		*res = append(*res, cur)
		return
	}
	if o > 0 {
		back2(res, cur+"(", o-1, c)
	}
	if c > o {
		back2(res, cur+")", o, c-1)
	}
}