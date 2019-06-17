package partition

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	before := &ListNode{}
	beforeHead := before
	after := &ListNode{}
	afterHead := after
	current := head
	for current != nil {
		if current.Val < x {
			before.Next = current
			before = before.Next
		} else {
			after.Next = current
			after = after.Next
		}
		current = current.Next
	}

	before.Next = afterHead.Next
	after.Next = nil

	return beforeHead.Next
}

