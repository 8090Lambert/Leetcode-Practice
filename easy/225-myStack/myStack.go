package _25_myStack

type MyStack struct {
	Tail *Node
	Length int
}

type Node struct {
	Value int
	Next *Node
	Prev *Node
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{nil, 0}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	newNode := &Node{x, nil, nil}
	if this.Tail == nil {
		this.Tail = newNode
	} else {
		this.Tail.Next = newNode
		newNode.Prev = this.Tail
		this.Tail = newNode
	}
	this.Length = this.Length + 1
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	newValue := this.Tail.Value
	if this.Tail.Prev != nil {
		this.Tail.Prev.Next = nil
		this.Tail = this.Tail.Prev
	} else {
		this.Tail = nil
	}
	
	if this.Length > 0 {
		this.Length = this.Length - 1
	}
	return newValue
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Tail == nil {
		return 0
	} else {
		return this.Tail.Value
	}
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Length == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */