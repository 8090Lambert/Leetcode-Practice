package _7_addBinary

import (
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