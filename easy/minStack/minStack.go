package minStack

//type MinStack struct {
//	MinData []int
//	Data []int
//}
//
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{}
//}
//
//
//func (this *MinStack) Push(x int)  {
//	if len(this.Data) == 0 {
//		this.Data = append(this.Data, x)
//		this.MinData = append(this.MinData, x)
//	} else {
//		if x <= this.GetMin() {
//			this.MinData = append(this.MinData, x)
//		}
//		this.Data = append(this.Data, x)
//	}
//}
//
//
//func (this *MinStack) Pop()  {
//	current := this.Data[len(this.Data)-1]
//	if current == this.GetMin() {
//		this.MinData = append([]int{}, this.MinData[0:len(this.MinData)-1]...)
//	}
//	this.Data = append([]int{}, this.Data[0:len(this.Data)-1]...)
//}
//
//
//func (this *MinStack) Top() int {
//	return this.Data[len(this.Data)-1]
//}
//
//
//func (this *MinStack) GetMin() int {
//	return this.MinData[len(this.MinData)-1]
//}




