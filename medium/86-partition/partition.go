package _6_partition

type ListNode struct {
	Val int
	Next *ListNode
}

func Partition(head *ListNode, x int) *ListNode {
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




func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var before, after *ListNode
	beforeHead := before
	afterHead := after
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	before.Next = afterHead.Next
	after.Next = nil
	return beforeHead.Next
}
