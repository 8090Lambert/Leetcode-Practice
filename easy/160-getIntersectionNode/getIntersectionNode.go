package _60_getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return b
}

func CountLength(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	lena, lenb := 0, 0
	for a.Next != nil {
		a = a.Next
		lena++
	}
	for b.Next != nil {
		b = b.Next
		lenb++
	}
	if a != b {
		return nil
	}
	a, b = headA, headB
	if lena > lenb {
		times := lena - lenb
		for times > 0 {
			a = a.Next
		}
	} else {
		times := lenb - lena
		for times > 0 {
			b = b.Next
		}
	}
	tmp := &ListNode{}
	for a != nil && b != nil {
		if a == b {
			tmp = b
			break
		}
		a = a.Next
		b = b.Next
	}

	return tmp
}
