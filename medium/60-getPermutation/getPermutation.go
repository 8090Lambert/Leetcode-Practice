package _0_getPermutation

import (
	"fmt"
	"strconv"
)

func GetPermutation(n int, k int) string {
	list := make([]int, n)
	for i := 1; i <= n; i++ {
		list[i-1] = i
	}
	res := ""
	count(&res, list, n, k-1)
	
	fmt.Println(res)
	return res
}

func count (res *string, list []int, n, k int) {
	if n == 0 {
		return
	}
	offset := k / factorial(n - 1)
	*res += strconv.Itoa(list[offset])
	temp := append([]int{}, list[:offset]...)
	if offset < len(list) - 1 {
		temp = append(temp, list[offset+1:]...)
	}
	count(res, temp, n-1, k % factorial(n-1))
}

func factorial (n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func getPermutation(n int, k int) string {
	factors := make([]int, n+1)
	factors[0] = 1
	numset := make([]int, n)
	for i := 1; i <= n; i++ {
		factors[i] = factors[i-1] * i
		numset[i-1] = i
	}
	k--
	var res string
	for n > 0 {
		n--
		a := k / factors[n]
		k %= factors[n]
		res = res + fmt.Sprintf("%d", numset[a])
		numset = append(numset[:a], numset[a+1:]...)
	}
	return res
}
