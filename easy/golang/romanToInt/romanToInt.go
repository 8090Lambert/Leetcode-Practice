package romanToInt

func romanToInt(s string) int {
	roman := map[byte]int{'I':1, 'V': 5, 'X': 10, 'L':50, 'C': 100, 'D': 500, 'M': 1000}
	length := len(s)
	total, prev, current := 0, 0, 0
	for i := 0; i < length; i++ {
		current = roman[s[i]]
		if current <= prev {
			total += current
		} else {
			total += current - 2 * prev
		}
		prev = current
	}

	return total
}