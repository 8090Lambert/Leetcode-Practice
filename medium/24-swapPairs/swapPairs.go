package _4_swapPairs

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

func SwapPairsLoop(head *ListNode) *ListNode {
	pre := &ListNode{Next:head}
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end
		start.Next = end.Next
		end.Next = start
		tmp = start
		//tmp = tmp.Next.Next
	}
	return pre.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := new(ListNode)
	pre.Next = head
	current, next := pre, head
	for next != nil && next.Next != nil {
		first, second := next, next.Next
		realNext := second.Next
		current.Next = second
		second.Next = first
		first.Next = realNext
		current = current.Next.Next
		next = realNext
	}
	if next != nil {
		current.Next = next
	}
	return pre.Next
}