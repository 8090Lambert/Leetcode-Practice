package _3_deleteDuplicates

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates (head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev, current := head, head.Next
	for current != nil {
		if prev.Val == current.Val {
			current = current.Next
		} else {
			prev.Next = current
			prev = prev.Next
			current = current.Next
		}
	}
	prev.Next = nil
	return head
}