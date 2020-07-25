package _8_rotate

func Rotate1(matrix [][]int)  {
	row := len(matrix)
	for i := 0; i < (row + 1) / 2; i++ {
		for j := 0; j < row / 2; j++ {
			tmp := matrix[row-1-j][i]
			matrix[row-1-j][i] = matrix[row-1-i][row-j-1]
			matrix[row-1-i][row-j-1] = matrix[j][row-1-i]
			matrix[j][row-1-i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}

func rotate(matrix [][]int) {
	row := len(matrix)
	for i := 0; i < (row+1)/2; i++ {
		for j := 0; j < row/2; j++ {
			tmp := matrix[row-1-j][i]
			matrix[row-1-j][i] = matrix[row-1-i][row-1-j]
			matrix[row-1-i][row-1-j] = matrix[j][row-1-i]
			matrix[j][row-1-i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}

