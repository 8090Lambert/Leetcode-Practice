package simpleStack

type MinStack struct {
	Contents []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.Contents) == 0 {
		this.Contents = []int{x, x}
	} else {
		currentMin := this.GetMin()
		if x < currentMin {
			this.Contents = append([]int{x, x}, this.Contents...)
		} else {
			this.Contents = append([]int{x, currentMin}, this.Contents...)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.Contents) != 0 {
		this.Contents = this.Contents[2:]
	}
}

func (this *MinStack) Top() int {
	return this.Contents[0]
}

func (this *MinStack) GetMin() int {
	if len(this.Contents) != 0 {
		return this.Contents[1]
	}

	return 0
}
