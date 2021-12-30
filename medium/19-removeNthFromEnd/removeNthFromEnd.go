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

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 1 && fast != nil {
		fast = fast.Next
		n--
	}
	if fast.Next == nil {
		return slow.Next
	}
	
	for fast.Next.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func RemoveNthFromEndOther(head *ListNode, n int) *ListNode {
	pre := new(ListNode)
	pre.Next = head
	cur := head
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	slow, fast := pre, cur
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
}

func Remove(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 0 && fast.Next != nil {
		fast = fast.Next
		n--
	}
	if n > 0 {
		return slow.Next
	}
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func removeNthFromEndOther(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	pre := dummy
	for cur.Next != nil {
		pre, cur = pre.Next, cur.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}