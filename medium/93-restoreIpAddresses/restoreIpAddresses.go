package _3_restoreIpAddresses

import "strconv"

func RestoreIpAddresses (s string) []string {
	collect := make([]string, 0)
	backTrack(s, 0, "", &collect)

	return collect
}


func backTrack (s string, start int, tmp string, res *[]string) {
	if start > 4 {
		return
	}

	if start == 4 && len(s) == 0 {
		*res = append(*res, tmp)
		return
	}

	for i := 1; i < 4; i++ {
		if len(s) < i {
			break
		}
		sub := s[:i]
		val, _ := strconv.Atoi(sub)
		if val > 255 || len(sub) > 1 && string(sub[0]) == "0" {
			continue
		}
		temp := tmp+sub
		if start != 3{
			temp += "."
		}

		backTrack(s[i:], start+1, temp, res)
	}
}



func restoreIpAddresses(s string) []string {
	collect := make([]string, 0)
	back(s, 0, "", &collect)
	return collect
}

func back(s string, start int, item string, res *[]string) {
	if start > 4 {
		return
	}
	if start == 0 && len(s) == 0 {
		*res = append(*res, item)
		return
	}
	for i := 1; i < 4; i++ {
		if len(s) < i {
			break
		}
		sub := s[:i]
		val, _ := strconv.Atoi(sub)
		if val > 255 || len(sub) > 0 && sub == "0" {
			continue
		}
		tmp := item + sub
		if start != 3 {
			tmp += "."
		}
		back(s[i:], start+1, tmp, res)
	}
}
