package __addTwoNumbers

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := new(ListNode)
	current := pre
	var one, two, over int
	for l1 != nil || l2 != nil {
		one, two = 0, 0
		if l1 != nil {
			one = l1.Val
		}
		if l2 != nil {
			two = l2.Val
		}
		sum := one + two + over
		if sum > 9 {
			over = 1
			sum = 0
		} else {
			over = 0
		}
		current.Next = &ListNode{Val:sum}
		current = current.Next
	}
	if over == 1 {
		current.Next = &ListNode{Val:1}
	}
	return pre.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := new(ListNode)
	current := pre
	over := 0
	for l1 != nil || l2 != nil {
		one, two := 0, 0
		if l1 != nil {
			one = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			two = l2.Val
			l2 = l2.Next
		}
		sum := one + two + over
		if sum > 9 {
			sum = sum % 10
			over = 1
		} else {
			over = 0
		}
		current.Next = &ListNode{Val:sum}
		current = current.Next
	}
	if over != 0 {
		current.Next = &ListNode{Val:1}
	}
	return pre.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	new := new(ListNode)
	cur := new
	plus := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + plus
		if sum > 9 {
			sum %= 10
			plus = 1
		} else {
			plus = 0
		}
		l1 = l1.Next
		l2 = l2.Next
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}
	for l1 != nil {
		sum := l1.Val + plus
		l1 = l1.Next
		if sum > 9 {
			plus = 1
			sum %= 10
		} else {
			plus = 0
		}
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}
	for l2 != nil {
		sum := l2.Val + plus
		l2 = l2.Next
		if sum > 9 {
			plus = 1
			sum %= 10
		} else {
			plus = 0
		}
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}
	if plus == 1 {
		cur.Next = &ListNode{Val: 1}
	}

	return new.Next
}