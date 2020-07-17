package _19_getRow

func GetRow(rowIndex int) []int {
	res := make([]int, 0)
	if rowIndex == 0 {
		return append(res, 1)
	}

	half := rowIndex / 2
	for index := 0; index <= rowIndex; index++ {
		if index == 0 {
			res = append(res, 1)
			continue
		}
		if index > half {
			res = append(res, res[rowIndex-index])
		} else {
			tmp := 1 * res[index-1] * (rowIndex - (index - 1)) / index
			res = append(res, tmp)
		}
	}

	return res
}

func getRow1(rowIndex int) []int {
	res := make([]int, 0, rowIndex+1)
	res = append(res, 1)
	for i := 1; i <= rowIndex; i++ {
		res = append(res, 1)
		for j := i-1; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}
