package getPermutation

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
	list := make([]int, n)
	for i := 1; i <= n; i ++ {
		list[i-1] = i
	}
	res := ""
	spell(list, &res, n, k-1)
	return res
}

func spell (list []int, res *string, n, k int) {
	if n == 0 {
		return
	}
	rank := rank(n-1)
	offset := k / n
	*res = strconv.Itoa(list[offset])
	temp := append([]int{}, list[:offset]...)
	if offset < len(list) - 1 {
		temp = append(temp, list[offset+1:]...)
	}
	
	spell(temp, res, n-1, k % rank)
}


func rank (n int) int {
	if n == 0 {
		return n
	}
	
	return n * rank(n-1)
}
