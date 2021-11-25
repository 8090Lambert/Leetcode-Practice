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
	res, cur := 1.0, x
	for ; N > 0; N /= 2 {
		if N % 2 == 1 {
			res *= cur
		}
		cur *= cur
	}
	return res
}