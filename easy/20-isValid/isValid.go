package _0_isValid

func isValid(s string) bool {
	count := len(s)
	if count & 1 == 1 {
		return false
	}
	match := map[byte]byte{
		')': '(', ']': '[', '}': '{',
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < count; i++ {
		if _, ok := match[s[i]]; !ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && match[s[i]] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}



func isValid1 (s string) bool {
	count := len(s)
	if count == 0 {
		return true
	} else if count & 1 == 1 {
		return false
	}
	match := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0, count)
	for i := 0; i < count; i++ {
		if need, ok := match[s[i]]; !ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == need {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}