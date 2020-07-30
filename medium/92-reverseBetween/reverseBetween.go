package _2_reverseBetween

type ListNode struct {
	Val int
	Next *ListNode
}

func ReverseBetween (head *ListNode, m, n int) *ListNode {
	if m == n || head == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy
	for i := 0; i < m - 1; i++ {
		prev = prev.Next
	}

	start := prev.Next
	tail := start.Next
	for i := 0; i < n - m; i++ {
		start.Next = tail.Next
		tail.Next = prev.Next
		prev.Next = tail
		tail = start.Next
	}

	return dummy.Next
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next:head}
	prev := dummy
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	start, tail := prev.Next, prev.Next.Next
	for i := 0; i < n-m; i++ {
		start.Next = tail.Next
		tail.Next = prev.Next
		prev.Next = tail
		tail = start.Next
	}
	return dummy.Next
}