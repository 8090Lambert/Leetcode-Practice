package isValid

func isValid(s string) bool {
	count := len(s)
	if count & 1  != 0 {
		return false
	}
	match := map[byte]byte{')':'(', '}': '{', ']': '['}
	stack := make([]byte, 0, count)
	for i := 0; i < count; i++{
		if pair, ok := match[s[i]]; ok {
			if len(stack) <= 0 || pair != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
