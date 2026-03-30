func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				dfs(i, j, row, col, &grid)
				count++
			}
		}
	}

	return count
}

func dfs(i, j, row, col int, grid *[][]byte) {
	if i < 0 || j < 0 || i >= row || j >= col || (*grid)[i][j] == '0' {
		return
	}

	(*grid)[i][j] = '0'

	dfs(i + 1, j, row, col, grid)
	dfs(i - 1, j, row, col, grid)
	dfs(i, j + 1, row, col, grid)
	dfs(i, j - 1, row, col, grid)
}
