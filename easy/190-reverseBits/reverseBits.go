package _90_reverseBits

func ReverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1
		res += num & 1
		num >>= 1
	}

	return res
}


func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32 && num > 0; i++ {
		res |= num & 1 << (31-i)
		num >>=1
	}
	return res
}
