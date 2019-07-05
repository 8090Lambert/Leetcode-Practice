package isUgly

func isUgly (num int) bool {
	if num == 0 {
		return false
	}
	
	for num != 1 {
		if num % 2 == 0 {
			num /= 2
		} else if num % 3 == 0 {
			num /= 3
		} else if num % 5 == 0 {
			num /=5
		} else {
			return false
		}
	}
	
	return true
}

func isUglyRecursion (num int) bool {
	if num == 2 || num == 3 || num == 5 || num == 1 {
		return true
	}
	
	if num != 0 {
		if num % 2 == 0 {
			return isUglyRecursion(num / 2)
		} else if num % 3 == 0 {
			return isUglyRecursion(num / 3)
		} else if num % 5 == 0 {
			return isUglyRecursion(num / 5)
		} else {
			return false
		}
	}
	
	return false
}
