package numArray

type NumArray struct {
	cell []int
}


func Constructor(nums []int) NumArray {
	return NumArray{
		nums,
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	res := 0
	if i < len(this.cell) && j < len(this.cell) {
		for k := i; k <= j; k++ {
			res += this.cell[k]
		}
	}

	return res
}
