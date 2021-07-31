package _1_mergeTwoLists

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	new := &ListNode{}
	current := new
	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			current.Next = l2
			l2 = l2.Next
		} else {
			current.Next = l1
			l1 = l1.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	} else if l2 != nil {
		current.Next = l2
	}
	return new.Next
}



func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	new := new(ListNode)
	current := new
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			current.Next = l2
			l2 = l2.Next
		} else {
			current.Next = l1
			l1 = l1.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return new
}
