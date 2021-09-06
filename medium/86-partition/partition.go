package _6_partition

type ListNode struct {
	Val int
	Next *ListNode
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

func Partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := new(ListNode), new(ListNode)
	l, r := left, right
	for head != nil {
		if head.Val < x {
			l.Next = head
			l = l.Next
		} else {
			r.Next = head
			r = r.Next
		}
		head = head.Next
	}
	l.Next = right.Next
	if r != nil {
		r.Next = nil
	}
	return left.Next
}