package josephring

type Node struct {
	Data int
	Next *Node
}

func JosephRing (m, n int) *Node {
	list := createCircleList(m)

	i := 1
	for list.Next != list {
		if (i + 1) % n == 0 {
			list.Next = list.Next.Next
			i = 1
		} else {
			i++
		}
		list = list.Next
	}

	return list
}

func createCircleList (m int) *Node {
	i := 1
	node := &Node{1,nil}
	head := node
	for i < m {
		i++
		node.Next = &Node{i, nil}
		node = node.Next
	}
	node.Next = head

	return node.Next
}
