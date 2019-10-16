package addTwoNumbers

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	current := pre
	overTen := 0
	for l1 != nil || l2 != nil {
		one, two := 0, 0
		if l1 != nil {
			one = l1.Val
		}
		if l2 != nil {
			two = l2.Val
		}
		sum := one+two+overTen
		if sum >= 10 {
			sum %= 10
			overTen = 1
		} else {
			overTen = 0
		}
		current.Next = &ListNode{Val:sum}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		current = current.Next
	}

	if overTen == 1 {
		current.Next = &ListNode{Val:1}
	}

	return pre.Next
}
