package main

import (
	reverseBits "./reverseBits"
)

func main() {
	var a uint32
	a = 43261596

	reverseBits.ReverseBits(a)
}
