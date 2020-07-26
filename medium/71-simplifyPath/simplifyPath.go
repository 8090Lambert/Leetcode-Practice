package _1_simplifyPath

import (
	"strings"
)

func SimplifyPath(path string) string {
	res := make([]string,0)
	arr := strings.Split(path, "/")
	for i := 1; i < len(arr); i++ {
		if arr[i] == "." || arr[i] == "" {
			continue
		} else if arr[i] == ".." {
			if len(res) - 1 < 0 {
				res = []string{}
			} else {
				res = append([]string{}, res[:len(res)-1]...)
			}
		} else {
			res = append(res, arr[i])
		}
	}

	return "/"+strings.Trim(strings.Join(res, "/"), "/")
}

func simplifyPath(path string) string {
	res := make([]string, 0)
	arr := strings.Split(path, "/")
	for i := 1; i < len(arr); i++ {
		if arr[i] == "." || arr[i] == "" {
			continue
		} else if arr[i] == ".." {
			count := len(res)
			if count == 0 {
				res = []string{}
			} else {
				res = append([]string{}, res[:count-1]...)
			}
		} else {
			res = append(res, string(arr[i]))
		}
	}
	return "/" + strings.Join(res, "/")
}