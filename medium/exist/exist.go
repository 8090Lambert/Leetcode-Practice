package exist

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	visited := make([][]bool, 0)
	for n := 0; n < row; n++ {
		item := make([]bool, col)
		visited = append(visited, item)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] && backTrack(i, j, 0, board, word, visited) {
				return true
			}
		}
	}

	return false
}

func backTrack (i, j, index int, board [][]byte, word string, visited [][]bool) bool {
	if index == len(word) {
		return true
	}
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[index] || visited[i][j] == true {
		return false
	}
	visited[i][j] = true

	if backTrack(i+1, j, index+1, board, word, visited) == true ||
		backTrack(i, j+1, index+1, board, word, visited) == true ||
		backTrack(i-1, j, index+1, board, word, visited) == true ||
		backTrack(i, j-1, index+1, board, word, visited) == true {
		return true
	}
	visited[i][j] = false

	return false
}