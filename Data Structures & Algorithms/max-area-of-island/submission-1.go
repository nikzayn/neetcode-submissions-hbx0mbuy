func maxAreaOfIsland(grid [][]int) int {
    row, col := len(grid), len(grid[0])
	maxCount := 0

    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if grid[i][j] == 1 {
                maxCount = max(maxCount, dfs(i, j, row, col, &grid))
            }
        }
    }

    return maxCount
}

func dfs(i, j, row, col int, grid *[][]int) int {
    if i < 0 || j < 0 || i >= row || j >= col || (*grid)[i][j] == 0 {
        return 0
    }

    (*grid)[i][j] = 0

    return 1 + dfs(i + 1, j, row, col, grid) +
    dfs(i - 1, j, row, col, grid) + 
    dfs(i, j + 1, row, col, grid) +
    dfs(i, j - 1, row, col, grid)
}
