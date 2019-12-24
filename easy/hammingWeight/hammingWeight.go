package hammingWeight

func HammingWeight(num uint32) int {
	var res uint32
	for num > 0 {
		res = res + (num & 1)
		num >>= 1
	}

	return int(res)
}
