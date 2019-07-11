package reverseBetween

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
	prev := head
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
