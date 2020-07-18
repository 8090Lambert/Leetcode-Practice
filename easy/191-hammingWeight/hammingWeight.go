package _91_hammingWeight

func HammingWeight(num uint32) int {
	var res uint32
	for num > 0 {
		res = res + (num & 1)
		num >>= 1
	}

	return int(res)
}

func hammingWeight(num uint32) int {
	num1 := int(num)
	res := 0
	for num1 > 0 {
		res += num1 & 1
		num1 >>= 1
	}
	return res
}
