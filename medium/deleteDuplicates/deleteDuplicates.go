package deleteDuplicates

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{}
	newHead.Next = head
	slow := newHead
	fast := head

	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			value := fast.Val
			for fast != nil && value == fast.Val {
				fast = fast.Next
			}
		} else {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}

	slow.Next = fast

	return newHead.Next
}
