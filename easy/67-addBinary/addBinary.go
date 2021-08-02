package _7_addBinary

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if a == "0" && b == "0" {
		return "0"
	}
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	} else {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}
	res, carry, length := "", false, len(a)-1
	for i := length; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if carry == true {
				res = "1" + res
			} else {
				res = "0" + res
			}
			carry = true
		} else if a[i] == '1' && b[i] == '0' || a[i] == '0' && b[i] == '1' {
			if carry == true {
				res = "0" + res
				carry = true
			} else {
				res = "1" + res
				carry = false
			}
		} else {
			if carry == true {
				res = "1" + res
			} else {
				res = "0" + res
			}
			carry = false
		}
	}
	if carry == true {
		res = "1" + res
	}
	return res
}


func addBinaryOther(a string, b string) string {
	lena := len(a)
	lenb := len(b)
	if lena - lenb < 0 {
		a = strings.Repeat("0", lenb-lena) + a
	} else {
		b = strings.Repeat("0", lena-lenb) + b
	}

	var plus uint8 = 0
	newStr := ""
	for i := len(a)-1; i >= 0; i-- {
		newStr = strconv.Itoa(int(a[i] ^ b[i] ^ plus)) + newStr
		if a[i] ^ b[i] == 0 {
			if a[i] == '1' && b[i] == '1' {
				plus = 1
			} else {
				plus = 0
			}
		} else {
			if plus == 1 {
				plus = 1
			} else {
				plus = 0
			}
		}
	}
	if plus == 1 {
		newStr = "1" + newStr
	}
	return newStr
}