package reverseBits

func ReverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; {
		res <<= 1
		res += num & 1
		num >>= 1
		i++
	}

	return res
}
