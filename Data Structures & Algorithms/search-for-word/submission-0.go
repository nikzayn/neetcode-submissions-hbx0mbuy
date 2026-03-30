func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	row, col := len(board), len(board[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(i, j, 0, row, col, &board, word) {
				return true
			}
		}
	}

	return false
}

func dfs(i, j, start, row, col int, board *[][]byte, word string) bool {
	if start == len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= row || j >= col || (*board)[i][j] != word[start] {
		return false
	}

	temp := (*board)[i][j]
	(*board)[i][j] = '#'

	found := dfs(i+1, j, start + 1, row, col, board, word) ||
	dfs(i-1, j, start + 1, row, col, board, word) ||
	dfs(i, j+1, start + 1, row, col, board, word) ||
	dfs(i, j-1, start + 1, row, col, board, word)

	(*board)[i][j] = temp
	return found
}