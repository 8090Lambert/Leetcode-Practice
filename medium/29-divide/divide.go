package _9_divide


func divide(dividend int, divisor int) int {
	position := -1
	if dividend < 0 {
		dividend *= -1
		position *= -1
	}
	if divisor < 0 {
		divisor *= -1
		position *= -1
	}
	var guess uint = 31
	var compare int = 0
}
