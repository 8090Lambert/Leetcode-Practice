package moveZeroes

func moveZeroes (nums []int) {
	i := 0
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[i], nums[cur] = nums[cur], nums[i]
			i++
		}
	}
}




func moveZeroes1(nums []int)  {
	var cur, next int = 0, 0
	for cur <= len(nums)-1 && next <= len(nums)-1 {
		if nums[next] != 0 {
			nums[cur], nums[next] = nums[next], nums[cur]
			cur++
		}
		next++
	}
}