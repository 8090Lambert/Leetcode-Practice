package _8_merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	sum := m + n
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[sum-1] = nums1[m-1]
			m--
		} else {
			nums1[sum-1] = nums2[n-1]
			n--
		}
		sum--
	}
	for i := n; i > 0; i-- {
		nums1[sum-1] = nums2[i-1]
		sum--
	}
}

func mergeArr(nums1 []int, m int, nums2 []int, n int) {
	sum := m + n
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[sum-1] = nums1[m-1]
			m--
		} else {
			nums1[sum-1] = nums2[n-1]
			n--
		}
		sum--
	}
	for i := n-1; i >= 0; i-- {
		nums1[sum-1] = nums2[n]
		sum--
	}
}
