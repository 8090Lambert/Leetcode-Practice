package intToRoman

func IntToRoman(num int) string {
	var romanString = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var roman = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var res string
	count := len(roman)
	for index := 0; index < count; index++ {
		for num >= roman[index] {
			num -= roman[index]
			res += romanString[index]
		}
	}

	return res
}
