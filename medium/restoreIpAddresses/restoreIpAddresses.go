package restoreIpAddresses

import "strconv"

func restoreIpAddresses (s string) []string {
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
		if start != 3 {
			temp = tmp+sub+"."
		}

		backTrack(s[i:], start+1, temp, res)
	}
}
