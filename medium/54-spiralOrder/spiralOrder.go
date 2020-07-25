package _4_spiralOrder

func SpiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) <= 0 {
		return res
	}

	rowStart, colStart := 0, 0
	rowEnd, colEnd := len(matrix)-1, len(matrix[0])-1

	for rowStart <= rowEnd && colStart <= colEnd {
		for c := colStart; c <= colEnd; c++ {
			res = append(res, matrix[rowStart][c])
		}
		for r := rowStart + 1; r <= rowEnd; r++{
			res = append(res, matrix[r][colEnd])
		}

		if rowStart < rowEnd && colStart < colEnd {
			for c := colEnd-1; c > colStart; c--{
				res = append(res, matrix[rowEnd][c])
			}

			for r := rowEnd; r > rowStart; r-- {
				res = append(res, matrix[r][colStart])
			}
		}

		colStart++
		colEnd--
		rowStart++
		rowEnd--
	}

	return res
}



func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) <= 0 {
		return res
	}
	rowStart, colStart := 0, 0
	rowEnd, colEnd := len(matrix)-1, len(matrix[0])-1
	for rowStart <= rowEnd && colStart <= colEnd {
		for c := colStart; c <= colEnd; c++ {
			res = append(res, matrix[rowStart][c])
		}
		for r := rowStart+1; r <= rowEnd; r++ {
			res = append(res, matrix[r][colEnd])
		}
		if rowStart < rowEnd && colStart < colEnd {
			for c := colEnd-1; c >= colStart; c-- {
				res = append(res, matrix[rowEnd][c])
			}
			for r := rowEnd-1; r >= rowStart+1; r-- {
				res = append(res, matrix[r][colStart])
			}
		}
		rowStart++
		colStart++
		rowEnd--
		colEnd--
	}
	return res
}