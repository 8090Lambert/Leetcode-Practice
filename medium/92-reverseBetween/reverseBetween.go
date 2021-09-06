package _2_reverseBetween

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right || head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur, next := pre.Next, pre.Next.Next
	for i := 0; i < right-left; i++ {
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
		next = cur.Next
	}
	return dummy.Next
}