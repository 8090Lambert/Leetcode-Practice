package getRow

func GetRow(rowIndex int) []int {
	res := make([]int, 0)
	if rowIndex == 0 {
		return append(res, 1)
	}
	if rowIndex == 1 {
		return append(res, 1, 1)
	}

	half := rowIndex / 2
	for i := 0; i < rowIndex+1; i++ {
		if i == 0 {
			res = append(res, 1)
			continue
		}
		if i > half { // 超过一半开始对称
			res = append(res, res[rowIndex-i])
		} else {
			tmp := 1 * res[i-1] * (rowIndex - (i - 1)) / i
			res = append(res, tmp)
		}
	}

	return res
}
