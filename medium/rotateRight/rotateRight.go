package rotateRight

type ListNode struct {
	Val int
	Next *ListNode
}

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var num int
	current := head
	for num = 0; current != nil; num++ {
		if current.Next == nil {
			current.Next = head
			current = nil
		} else {
			current = current.Next
		}
	}
	
	newTail := head
	for i := 0; i < num - k % num - 1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	
	return newHead
}
