package intersection

func intersection (nums1 []int, nums2 []int) []int {
	res := []int{}
	l1, l2 := len(nums1), len(nums2)
	if l2 < l1 {
		l1, l2 = l2, l1
		nums1, nums2 = nums2, nums1
	}
	intersectionMap := map[int]struct{}{}

	for i := 0; i < l1; i++ {
		if _, ok := intersectionMap[nums1[i]]; ok {
			continue
		} else {
			for j := 0; j < l2; j++ {
				if nums2[j] == nums1[i] {
					intersectionMap[nums1[i]] = struct{}{}
					res = append(res, nums1[i])
					break
				}
			}
		}
	}

	return res
}