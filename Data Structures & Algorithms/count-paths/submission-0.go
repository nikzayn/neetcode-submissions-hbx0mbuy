func uniquePaths(m int, n int) int {
    if m > n {
		m, n = n, m
	}

	res := 1

	for i := 1; i < m; i++ {
		res = res * (n + i - 1) / i
	}

	return res
}
