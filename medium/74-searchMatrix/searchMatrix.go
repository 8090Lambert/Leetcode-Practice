package _4_searchMatrix

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	left, right := 0, row * col - 1
	for left <= right {
		idx := (left + right) / 2
		if target == matrix[idx/col][idx%col] {
			return true
		} else if target < matrix[idx/col][idx%col] {
			right = idx - 1
		} else {
			left = idx + 1
		}
	}

	return false
}


/**
	从右上角（左下角）开始查找
 */
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	r, c := 0, col-1
	for r < row && c >= 0 {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] < target {
			r--
		} else {
			c++
		}
	}
	return false
}