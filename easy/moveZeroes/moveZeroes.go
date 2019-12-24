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
