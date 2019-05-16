package quick

func Quick(arr []int) []int {
	count := len(arr)
	if count <= 1 {
		return arr
	}
	start := arr[0]

	big := make([]int, 0)
	small := make([]int, 0)
	for index := 1; index < count; index++ {
		if arr[index] <= start {
			small = append(small, arr[index])
		} else {
			big = append(big, arr[index])
		}
	}

	small = quick(small)
	big = quick(big)

	return append(append(append([]int{}, big...), start), small...)
}

