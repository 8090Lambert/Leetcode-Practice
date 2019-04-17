package generate

func Generate(numRows int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		tmp := make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tmp = append(tmp, 1)
			} else {
				need := res[i-1][j-1] + res[i-1][j]
				tmp = append(tmp, need)
			}
		}
		res = append(res, tmp)
	}

	return res
}

func Generate1(numRows int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		tmp := make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tmp = append(tmp, 1)
			} else {
				need := res[i-1][j] + res[i-1][j-1]
				tmp = append(tmp, need)
			}
		}
		res = append(res, tmp)
	}

	return res
}
