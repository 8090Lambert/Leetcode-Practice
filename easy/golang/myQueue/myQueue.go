package myQueue

type MyQueue struct {
	Stack0 []int
	Stack1 []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.Stack0 = append(this.Stack0, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	need := this.Peek()
	this.Stack0 = append([]int{}, this.Stack0[1:]...)
	
	return need
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	for i := len(this.Stack0)-1; i >= 0; i-- {
		this.Stack1 = append(this.Stack1, this.Stack0[i])
	}
	need := this.Stack1[len(this.Stack1)-1]
	this.Stack1 = []int{}
	
	return need
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Stack0) == 0
}
