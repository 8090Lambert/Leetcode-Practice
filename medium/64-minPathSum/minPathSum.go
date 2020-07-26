package _4_minPathSum

import "math"

func MinPathSum(grid [][]int) int {
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

func minPathSummer (grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j >= 1 {
				grid[0][j] += grid[0][j-1]
			} else if i >= 1 && j == 0 {
				grid[i][0] += grid[i-1][0]
			} else if i >= 1 && j >= 1{
				grid[i][j] = int(math.Min(float64(grid[i-1][j]), float64(grid[i][j-1]))) + grid[i][j]
			}
		}
	}
	return grid[row-1][col-1]
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	for r := 1; r < row; r++ {
		grid[r][0] += grid[r-1][0]
	}
	for c := 1; c < col; c++ {
		grid[0][c] += grid[0][c-1]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			grid[i][j] = int(math.Min(float64(grid[i-1][j]), float64(grid[i][j-1]))) + grid[i][j]
		}
	}
	return grid[row-1][col-1]
}