package mergeKLists

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	
	return merge(lists, 0, len(lists))
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	
	mid := (start + end) / 2
	node1 := merge(lists, start, mid)
	node2 := merge(lists, mid + 1, end)
	
	return mergeTwoLists(node1, node2)
}

func mergeTwoLists (n1, n2 *ListNode) *ListNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	
	if n1.Val < n2.Val {
		n1.Next = mergeTwoLists(n1.Next, n2)
		return n1
	} else {
		n2.Next = mergeTwoLists(n1, n2.Next)
		return n2
	}
}


func mergeKList (lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	
	return merge1(lists, 0, len(lists)-1)
}

func merge1(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	
	mid := (start + end) / 2
	nodeLeft := merge1(lists, start, mid)
	nodeRight := merge1(lists, mid + 1, end)
	
	return mergeTwo(nodeLeft, nodeRight)
}


func mergeTwo (left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	
	if left.Val < right.Val {
		left.Next = mergeTwo(left.Next, right)
		return left
	} else {
		right.Next = mergeTwo(left, right.Next)
		return right
	}
}
