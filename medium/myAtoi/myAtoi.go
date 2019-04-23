package myAtoi

import (
	"fmt"
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
