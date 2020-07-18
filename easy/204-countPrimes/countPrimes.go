package _04_countPrimes

func countPrimes (n int) int {
	sum := make([]bool, n)
	for i := 2; i < n; i++ {
		if sum[i] == false {
			for j := 2; i * j < n; j++ {
				sum[i*j] = true
			}
		}
	}
	count := 0
	for i := 2; i < len(sum); i++ {
		if sum[i] == false {
			count += 1
		}
	}
	
	return count
}


func countPrimes2 (n int) int {
	sum := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if sum[i] == true {
			continue
		}
		for j := i + i; j < n; j+=i {
			sum[j] = true
		}
		count++
	}
	return count
}


func countPrimes1 (n int) int {
	sum := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if sum[i] {
			continue
		}
		for j := i; j < n; j+=i {
			sum[j] = true
		}
		count++
	}
	return count
}
