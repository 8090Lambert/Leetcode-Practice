package reverseList

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList (head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	
	init := &ListNode{}
	prev := init
	current := head
	for current != nil {
		if prev == init {
			prev = nil
		}
		originNext := current.Next
		current.Next = prev
		prev = current
		if originNext != nil {
			current = originNext
		} else {
			break
		}
	}
	
	return current
}
