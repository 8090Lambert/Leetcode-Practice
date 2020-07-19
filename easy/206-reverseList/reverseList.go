package _06_reverseList

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList (head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, next *ListNode
	next = head
	for next != nil {
		tmp := next.Next
		next.Next = prev
		prev = next
		next = tmp
	}

	return prev
}
