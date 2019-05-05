package swapPairs

type ListNode struct {
	Val int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	list := head.Next
	head.Next = SwapPairs(list.Next)
	list.Next = head

	return list
}
