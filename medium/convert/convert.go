package convert

import "strings"

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