package addBinary

import (
	"fmt"
	"strings"
)

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

		fmt.Println(result, carry)
	}
	//for intA != 0 || intB != 0 {
	//	if index >= len(a) && index >= len(b) {
	//		break
	//	}
	//	if intA % 10 == 0 && intB % 10 == 1 ||
	//	intA % 10 == 1 && intB % 10 == 0 {
	//		if jinwei == true {
	//			result = "0" + result
	//			jinwei = true
	//		} else {
	//			result = "1" + result
	//			jinwei = false
	//		}
	//	} else if intA % 10 == 0 && intB % 10 == 0 {
	//		if jinwei == true {
	//			result = "1" + result
	//		} else {
	//			result = "0" + result
	//		}
	//		jinwei = false
	//	} else {
	//		if jinwei == true {
	//			result = "1" + result
	//		} else {
	//			result = "0" + result
	//		}
	//		jinwei = true
	//	}
	//
	//	fmt.Println(result, jinwei)
	//	intA = intA / 10
	//	intB = intB / 10
	//	index++
	//}
	//
	if carry == true {
		result = "1" + result
	}

	return result
}

/*func AddBinary(a string, b string) string {
	aLen, bLen := len(a), len(b)
	max, min := aLen, bLen
	maxString, minString := a, b
	if max - min < 0 {
		max, min = bLen, aLen
		maxString, minString = b, a
	}
	jinwei := false

	asciiCodeMap := map[int]string{
		48: "0",
		49: "1",
	}
	var result string

	diff := max - min
	for i := max - 1; i >= 0 ; i-- {
		if i - diff < 0 {
			break
		}
		_, aok := asciiCodeMap[int(maxString[i])]
		_, bok := asciiCodeMap[int(minString[i - diff])]

		if !aok || !bok {
			return result
		}

		if asciiCodeMap[int(maxString[i])] == "1" && asciiCodeMap[int(minString[i - diff])] == "0" ||
			asciiCodeMap[int(maxString[i])] == "0" && asciiCodeMap[int(minString[i - diff])] == "1" {
			if jinwei == true {
				result = result + "0"
				jinwei = true
			} else {
				result = result + "1"
				jinwei = false
			}
		} else if asciiCodeMap[int(maxString[i])] == "1" && asciiCodeMap[int(minString[i - diff])] == "1"{
			if jinwei == true {
				result = result + "1"
			} else {
				result = result + "0"
			}
			jinwei = true
		} else {
			if jinwei == true {
				result = result + "1"
			} else {
				result = result + "0"
			}
			jinwei = false
		}
	}

	if jinwei == true {
		for i := diff - 1; i >= 0; i-- {
			if jinwei == false {
				break
			}
			if asciiCodeMap[int(maxString[i])] == "1" {
				result = result + "0"
				jinwei = true
			} else {
				result = result + "1"
				jinwei = false
			}
		}
	} else {
		result = string([]byte(maxString)[0:diff]) + result
	}

	if jinwei == true {
		result = "1" + result
	}

	return string(result)
}*/
