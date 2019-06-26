package removeElements

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements (head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dynamic := &ListNode{}
	dynamic.Next = head
	current := dynamic
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	
	return dynamic.Next
}