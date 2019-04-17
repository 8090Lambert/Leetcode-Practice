package merge

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	sum := m + n
	for m > 0 && n > 0{
		if nums1[m-1] >= nums2[n-1] {
			nums1[sum-1] = nums1[m-1]
			sum--
			m--
		} else {
			nums1[sum-1] = nums2[n-1]
			sum--
			n--
		}
	}

	for n > 0 {
		nums1[sum-1] = nums2[n-1]
		sum--
		n--
	}

	return nums1
}
