package isPalindrome

import (
	"strconv"
	"strings"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	origin := x
	var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	return result == origin
}

func IsPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	begin, length := 0, len(strconv.Itoa(x))
	res, count := 0, length / 2
	for begin < count {
		res = res * 10 + x % 10
		x /= 10
		begin++
	}
	if length & 1 == 1 {
		x /= 10
	}

	return x == res
}

// 回文字符串
func IsPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}
	strArr := strings.Split(strings.ToLower(s), "")
	start := 0
	end := len(strArr) - 1
	for start < end {
		for isValid(strArr[start]) == false && start < end {
			start++
		}
		for isValid(strArr[end]) == false && start < end {
			end--
		}
		if strArr[start] == strArr[end] {
			start++
			end--
			continue
		} else {
			return false
		}
	}

	return true
}

func isValid(s string) bool {
	return (s >= "a" && s <= "z") || (s >= "0" && s <= "9")
}

type ListNode struct {
     Val int
     Next *ListNode
}

func isPalindromeList (head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	
	fast := head
	slow := head
	stack := make([]*ListNode, 0)
	for fast.Next != nil && fast.Next.Next != nil {
		stack = append(stack, slow)
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	if fast.Next != nil {
		stack = append(stack, slow)
	}
	slow = slow.Next
	
	mid := slow
	for i := len(stack) - 1; i >= 0; i-- {
		if mid.Val != stack[i].Val {
			return false
		}
		mid = mid.Next
	}
	
	return true
}