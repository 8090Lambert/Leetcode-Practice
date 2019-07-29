package searchMatrix

func searchMatrix(matrix [][]int, target int) bool {
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