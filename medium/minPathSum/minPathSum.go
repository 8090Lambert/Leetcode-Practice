package minPathSum

import "math"

func minPathSum(grid [][]int) int {
	row, col := len(grid)-1, len(grid[0])-1

	for i := row; i >= 0; i-- {
		for j := col; j >= 0; j-- {
			if i == len(grid)-1 && j != len(grid[0]) - 1 {
				grid[i][j] = grid[i][j] + grid[i][j+1]
			} else if j == len(grid[0]) - 1 && i != len(grid) - 1 {
				grid[i][j] = grid[i][j] + grid[i+1][j]
			} else if j != len(grid[0]) - 1 && i != len(grid) - 1 {
				grid[i][j] = grid[i][j] + int(math.Min(float64(grid[i+1][j]), float64(grid[i][j+1])))
			}
		}
	}

	return grid[0][0]
}