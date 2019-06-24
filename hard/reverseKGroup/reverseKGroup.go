package reverseKGroup

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	current := head
	count := 0
	for current != nil && count != k {
		current = current.Next
		count++
	}

	if count == k {
		current := reverseKGroup(current, k)
		for count != 0 {
			tmp := head.Next
			head.Next = current
			current = head
			head = tmp
			count--
		}
		head = current
	}

	return head
}