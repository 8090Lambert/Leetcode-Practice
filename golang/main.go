package main

import (
	rotate "./rotate"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate.Rotate(a, 3)
}
