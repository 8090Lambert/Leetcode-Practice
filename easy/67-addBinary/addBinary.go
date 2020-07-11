package _7_addBinary

import (
	"strings"
)

func addBinary(a string, b string) string {
	if a == "0" && b == "0" {
		return "0"
	}
	diff := len(a) - len(b)
	if diff > 0 {
		b = strings.Repeat("0", diff) + b
	} else {
		a = strings.Repeat("0", diff*-1) + a
	}
	res := make([]byte, len(a), len(a))
	carry := 0
	for i := len(a)-1; i >= 0; i-- {
		if a[i] == 49 && b[i] == 49 {
			if carry == 1 {
				res[i] = 49
			} else {
				res[i] = 48
			}
			carry = 1
		} else if a[i] == 49 && b[i] == 48 || a[i] == 48 && b[i] == 49 {
			if carry == 1 {
				carry = 1
				res[i] = 48
			} else {
				carry = 0
				res[i] = 49
			}
		} else {
			if carry == 1 {
				res[i] = 49
			} else {
				res[i] = 48
			}
			carry = 0
		}
	}

	if carry == 1 {
		res = append([]byte{49}, res...)
	}

	return string(res)
}

func AddBinary(a string, b string) string {
	var result string
	if a == "0" && b == "0" {
		return "0"
	}
	diff := len(a) - len(b)
	short := &b
	if diff < 0 {
		diff = len(b) - len(a)
		short = &a
	}
	*(short) = strings.Repeat("0", diff) + *(short)

	carry := false
	length := len(a)
	for i := length - 1; i >= 0; i-- {
		if string(a[i]) == "0" && string(b[i]) == "1" ||
			string(a[i]) == "1" && string(b[i]) == "0" {
			if carry == true {
				result = "0" + result
				carry = true
			} else {
				result = "1" + result
				carry = false
			}
		} else if string(a[i]) == "0" && string(b[i]) == "0" {
			if carry == true {
				result = "1" + result
			} else {
				result = "0" + result
			}
			carry = false
		} else {
			if carry == true {
				result = "1" + result
			} else {
				result = "0" + result
			}
			carry = true
		}
	}
	if carry == true {
		result = "1" + result
	}

	return result
}

















