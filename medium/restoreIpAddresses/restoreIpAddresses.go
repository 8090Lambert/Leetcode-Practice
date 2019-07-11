package restoreIpAddresses

import (
	"strconv"
)

func RestoreIpAddresses (s string) []string {
	res := make([]string, 0)
	if len(s) == 0 || len(s)>12{
		return res
	}
	
	backTrack(s, 0, "", &res)
	return res
}


func backTrack (s string, n int, tmp string, res *[]string) {
	if n > 4 {
		return
	}
	
	if n == 4 && len(s) == 0 {
		*res = append(*res, tmp)
		return
	}
	
	for i := 1; i < 4; i++ {
		if len(s) < i {
			break
		}
		sub := s[0:i]
		val, _ := strconv.Atoi(sub)
		if val > 255 || len(sub) > 1 && sub[0] == '0' {
			continue
		}
		
		temp := tmp+sub
		if n < 3 {
			temp = tmp+sub+"."
		}
		
		backTrack(s[i:], n+1, temp, res)
	}
}
