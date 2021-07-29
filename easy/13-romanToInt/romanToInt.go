package _3_romanToInt

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res, count := 0, len(s)
	for i := range s {
		current := m[s[i]]
		if i < count-1 && current < m[s[i+1]] {
			res -= current
		} else {
			res += current
		}
	}
	return res
}


func romanToInt1(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res, prev := 0, 0
	for i := 0; i < len(s); i++ {
		res += m[s[i]]
		if m[s[i]] > prev {
			res -= 2 * prev
		}
		prev = m[s[i]]
	}
	return res
}