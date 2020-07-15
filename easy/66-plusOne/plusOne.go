package _6_plusOne

func PlusOne(digits []int) []int {
	length := len(digits)
	carry := 0
	if digits[length - 1] + 1 > 9 {
		carry = 1
		digits[length - 1] = 0
		for i := len(digits)-2; i >= 0; i-- {
			sum := digits[i] + 1
			if sum > 9 {
				digits[i] = 0
				carry = 1
			} else {
				digits[i] = sum
				carry = 0
				break
			}
		}
	} else {
		digits[length - 1] = digits[length - 1] + 1
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	length := len(digits)
	plus := 0
	for i := length - 1; i >= 0; i-- {
		sum := digits[i] + plus
		if plus == 0 {
			if i != length - 1 {
				break
			}
			sum += 1
		}
		if sum == 10 {
			plus = 1
		} else {
			plus = 0
		}
		digits[i] = sum % 10
	}

	if plus == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
