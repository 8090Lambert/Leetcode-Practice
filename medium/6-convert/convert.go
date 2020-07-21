package __convert

import "strings"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	b := []byte(s)
	n := len(s)
	cyc := 2 * numRows - 2
	res := make([]string, numRows)
	for i := 0; i < n; i++ {
		var mod = i % cyc
		if mod < numRows {
			res[mod] += string(b[i])
		} else {
			res[cyc-mod] += string(b[i])
		}
	}
	return strings.Join(res, "")
}

func Convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	n := len(s)
	cyclen := 2 * numRows - 2
	ret := make([]string, 0)
	for index :=0; index < numRows; index++ {
		for j := 0; j + index < n; j+=cyclen  {
			ret = append(ret, string(s[j+index]))
			if index != 0 && index + 1 != numRows && j + cyclen - 1 < n {
				ret = append(ret, string(s[j+cyclen-index]))
			}
		}
	}
	
	return strings.Join(ret, "")
}









func convert1(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	cyc := numRows * 2 - 2
	count := len(s)
	res := make([]string, numRows)
	for i := 0; i < count; i++ {
		var mod = i % cyc
		if mod < numRows {
			res[mod] += string(s[i])
		} else {
			res[cyc-mod] += string(s[i])
		}
	}
	return strings.Join(res, "")
}