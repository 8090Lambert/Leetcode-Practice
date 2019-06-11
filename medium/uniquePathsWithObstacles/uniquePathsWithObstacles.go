package uniquePathsWithObstacles

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	obstacleGrid[0][0] = 1
	for r := 1; r < row; r++ {
		if obstacleGrid[r][0] == 0 && obstacleGrid[r-1][0] == 1 {
			obstacleGrid[r][0] = 1
		} else {
			obstacleGrid[r][0] = 0
		}
	}

	for c := 1; c < col; c++ {
		if obstacleGrid[0][c] == 0 && obstacleGrid[0][c-1] == 1 {
			obstacleGrid[0][c] = 1
		} else {
			obstacleGrid[0][c] = 0
		}
	}


	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}

	return obstacleGrid[row-1][col-1]
}
