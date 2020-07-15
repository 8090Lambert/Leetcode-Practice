package __isPalindrome

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	length := len(strconv.Itoa(x))
	res, count := 0, length / 2
	for i := 0; i < count; i++ {
		res = res * 10 + x % 10
		x /= 10
	}
	if length & 1 == 1 {	// 奇数长度的中间位不要
		return res == (x / 10)
	} else {
		return res == x
	}
}

// 回文字符串
func isPalindromeString(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	count := len(s)
	start, end := 0, count - 1
	for start <= end {
		if !isValidString(s[start]) {
			start++
			continue
		} else if !isValidString(s[end]) {
			end--
			continue
		}
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isValidString(cell uint8) bool {
	b := byte(cell)
	if b >= 'a' && b <= 'z' || b >= '0' && b <= '9' {
		return true
	}
	return false
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