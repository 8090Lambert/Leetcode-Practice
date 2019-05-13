package main

import (
	"fmt"

	"leetcode/medium/nextPermutation"
)

func main() {
	fmt.Println("begin")
	a := []int{1, 3, 4, 56}
	nextPermutation.NextPermutation(a)
	fmt.Println("end")
}
