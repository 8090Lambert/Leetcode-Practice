package _3_romanToInt


func romanToInt(s string) int {
	mapper := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res, prev, current := 0, 0, 0
	for i := 0; i < len(s); i++ {
		current = mapper[s[i]]
		if current > prev {
			res += current
		} else {
			res = res + current - prev * 2
		}
		prev = current
	}
	return res
}