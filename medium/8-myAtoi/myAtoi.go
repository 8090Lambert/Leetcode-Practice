package __myAtoi

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MyAtoi(str string) int {
	str = strings.Trim(str, " ")
	count := len(str)
	res := 0
	flag := 0
	if string(str[0]) == "-" {
		flag = -1
	}  else if string(str[0]) == "+" || isValid(string(str[0])){
		flag = 1
	}
	
	min, max := -2147483648, 2147483647
	for index := 0; index < count; index++ {
		if isValid(string(str[index])) {
			tmp, _ := strconv.Atoi(string(str[index]))
			if res > (max - tmp) / 10 {
				if flag == 1 {
					return max
				} else {
					return min
				}
			}
			res = res * 10 + tmp
		} else {
			if index == 0{
				if string(str[index]) != "-" && string(str[index]) != "+" {
					return 0
				}
			} else {
				if res != 0 {
					break;
				}
			}
		}
	}
	fmt.Println(flag)
	return res * flag
}

func isValid(s string) bool {
	if s >= "0" && s <= "9" {
		return true
	}
	return false
}


func myAtoi(str string) int {
	min, max := int(math.Pow(-2, 31))-1, int(math.Pow(2, 31)) - 1
	str = strings.TrimSpace(str)
	count := len(str)
	res, flag := 0, 1
	if str[0] == '-' {
		flag = -1
	}
	for i := 0; i < count; i++ {
		if str[i] < '0' || str[i] > '9' {
			if i == 0 && (str[i] == '-' || str[i] == '+') {
				continue
			} else {
				break
			}
		} else {
			int, _ := strconv.Atoi(string(str[i]))
			if res > (max - int) / 10 {
				if flag == 1 {
					return max
				} else {
					return min
				}
			}
			res = res * 10 + int
		}
	}

	return res * flag
}