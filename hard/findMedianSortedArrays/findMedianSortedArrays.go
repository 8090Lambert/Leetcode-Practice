package findMedianSortedArrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	
	imin := 0
	imax := m
	halfLen := (m+n+1) / 2
	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i
		if i < imax && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > imin && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}
			if (m + n) % 2 == 1 {
				return float64(maxLeft)
			}
			
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = int(math.Min(float64(nums2[j]), float64(nums1[i])))
			}
			return float64((maxLeft + minRight) / 2)
		}
	}
	return 0.0
}
