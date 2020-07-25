package _0_myPow

func MyPow(x float64, n int) float64 {
	N := n
	if N < 0 {
		x = 1/x
		N = -N
	}

	var res float64 = 1.0
	var current float64 = x
	for i := N; i > 0; i /= 2 {
		if i % 2 == 1 {
			res = res * current
		}

		current = current * current
	}

	return res
}

func myPow(x float64, n int) float64 {
	N := n
	if N < 0 {
		x = 1/x
		N = -N
	}
	res := 1.0
	cur := x
	for i := N; i > 0; i /= 2 {
		if i % 2 == 1 {
			res = res * cur
		}
		cur *= cur
	}
	return res
}