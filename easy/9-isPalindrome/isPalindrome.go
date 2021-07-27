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

func isValidString(b uint8) bool {
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


func IsPalindromeList1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var pre, cur *ListNode
	cur = slow
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	mid := pre
	for mid != nil {
		if mid.Val != head.Val {
			return false
		}
		mid = mid.Next
		head = head.Next
	}
	return true
}


func isPalindromeNum(x int) bool {
	if x < 0 {
		return false
	}
	l := 0
	xc := x
	for xc > 0 {
		xc /= 10
		l++
	}
	res := 0
	for i := 0; i < (l/2); i++ {
		res = (res * 10) + (x % 10)
		x /= 10
	}
	if l & 1 == 1 {
		return res == (x/10)
	} else {
		return x == res
	}
}