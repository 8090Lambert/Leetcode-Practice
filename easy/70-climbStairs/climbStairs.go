package _0_climbStairs

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return ClimbStairs(n-1) + ClimbStairs(n - 2)
	}
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	first, second, sum := 1, 2, 0
	for i := 3; i <= n; i++ {
		sum = first + second
		first = second
		second = sum
	}
	return sum
}

func climbStairsRecursion(n int) int {
	first, second := 1, 1
	return climb(n, first, second)
}

func climb(n, f1, f2 int) int {
	if n <= 1 {
		return f2
	}
	return climb(n-1, f2, f1+f2)
}