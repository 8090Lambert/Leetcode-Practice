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

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{Next:head}
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end
		start.Next = end.Next
		end.Next = start
		tmp = start
	}
}


func test (head *ListNode) *ListNode {
	pre := &ListNode{Next:head}
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end
		start.Next = end.Next
		end.Next = start
		tmp = start
	}

	return pre.Next
}
