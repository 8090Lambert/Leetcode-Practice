package minStack

type MinStack struct {
	Data []int
	L int
}


/** initialize your Data structure here. */
func Constructor() MinStack {
	return MinStack{
		Data: make([]int, 0),
		L: -1,
	}
}


func (this *MinStack) Push(val int)  {
	if this.L > -1 {
		min := this.Data[this.L]
		if val < min {
			min = val
		}
		this.Data = append(this.Data, val, min)
	} else {
		this.Data = append(this.Data, val, val)
	}
	this.L += 2
}


func (this *MinStack) Pop() {
	//needReturn := this.Data[this.L-1]
	this.Data = append([]int{}, this.Data[:this.L-1]...)
	if this.L > -1 {
		this.L -= 2
	}
	//this.Data = append([]int{}, this.Data[:this.L]...)
}


func (this *MinStack) Top() int {
	return this.Data[this.L-1]
}


func (this *MinStack) GetMin() int {
	return this.Data[this.L]
}
