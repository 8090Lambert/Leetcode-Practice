package testStack

type MyStack struct {
	Data []int
}


/** Initialize your data structure here. */
func Constructor1() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.Data = append(this.Data, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	cur := this.Data[len(this.Data)-1]
	this.Data = append([]int{}, this.Data[:len(this.Data)-1]...)
	return cur
}


/** Get the top element. */
func (this *MyStack) Top() int {
	cur := this.Data[len(this.Data)-1]
	return cur
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Data) == 0
}
