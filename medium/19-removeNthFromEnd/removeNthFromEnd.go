package _9_removeNthFromEnd

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	left := current
	right := current
	for n > 1 && right.Next != nil{
		right = right.Next
		n--
	}
	if n > 1 {
		return nil
	}
	// 如果right刚好是尾节点，直接删除头节点
	if right.Next == nil {
		return left.Next
	}

	for right.Next.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next
	return current
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	for n > 1 && fast != nil {
		fast = fast.Next
		n--
	}
	if fast.Next == nil {
		return slow.Next
	}
	for fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}