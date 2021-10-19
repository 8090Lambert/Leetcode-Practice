package _2_deleteDuplicates

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{Next:head}
	slow := pre
	fast := head
	for fast != nil && fast.Next != nil {
		if fast.Val == fast.Next.Val {
			value := fast.Val
			for fast != nil && fast.Val == value {
				fast = fast.Next
			}
		} else {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}
	slow.Next = fast
	return pre.Next
}



func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := new(ListNode)
	pre.Next = head
	slow, fast := pre, head
	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			v := fast.Val
			for fast != nil && fast.Val == v {
				fast = fast.Next
			}
		} else {
			slow.Next = fast
			slow = slow.Next
			fast = fast.Next
		}
	}
	slow.Next = fast
	return pre.Next
}
