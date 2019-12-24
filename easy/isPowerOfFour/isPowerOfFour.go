package isPowerOfFour

func isPowerOfFour (num int) bool {
	return num > 0 && num&(num-1) == 0 && num & 0x55555555 == num
}
