package containsNearbyDuplicate

func containsNearbyDuplicate (nums []int, k int) bool {
	hasMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := hasMap[nums[i]]; ok && i - j <= k {
			return true
		}
		hasMap[nums[i]] = i
	}
	
	return false
}