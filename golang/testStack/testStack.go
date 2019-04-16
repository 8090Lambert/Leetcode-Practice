package testStack

import "math"

type MinStack struct {
	Min      int
	Contents []int
	Length   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Contents = append([]int{x}, this.Contents...)
	if this.Length == 0 {
		this.Min = x
	} else {
		if x < this.Min {
			this.Min = x
		}
	}

	this.Length = this.Length + 1
}

func (this *MinStack) Pop() {
	if this.Length > 1 {
		if this.Contents[0] <= this.Min {
			min := this.Contents[1]
			for index := this.Length - 1; index > 1; index-- {
				min = int(math.Min(float64(min), float64(this.Contents[index])))
			}
			this.Min = min
		}
		this.Length = this.Length - 1
		this.Contents = this.Contents[1:]
	} else {
		this.Contents = []int{}
		this.Length = 0
		this.Min = 0
	}
}

func (this *MinStack) Top() int {
	return this.Contents[0]
}

func (this *MinStack) GetMin() int {
	return this.Min
}
