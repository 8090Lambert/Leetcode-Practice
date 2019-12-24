package containsDuplicate

func containsDuplicate (nums []int) bool {
	unique := map[int]struct{}{}
	for i := 0; i < len(nums); i ++ {
		if _, ok := unique[nums[i]]; !ok {
			unique[nums[i]] = struct{}{}
		} else {
			return true
		}
	}
	
	return false
}
