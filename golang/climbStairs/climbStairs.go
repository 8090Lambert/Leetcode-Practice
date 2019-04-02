package climbStairs

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return ClimbStairs(n-1) + ClimbStairs(n - 2)
	}
}

func ClimbStairs1(n int) int {
	var third int
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		first, second := 1, 2
		for i := 3; i <= n; i++ {
			third = first + second
			first = second
			second = third
		}
	}

	return third
}