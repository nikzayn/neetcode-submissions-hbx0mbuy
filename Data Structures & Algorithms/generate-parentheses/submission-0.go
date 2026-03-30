var current = make([]byte, 16)

func generateParenthesis(n int) []string {
	var res []string

	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left + right == n*2 {
			res = append(res, string(current[:n*2]))
		}

		if left < n {
			current[left+right] = '('
			dfs(left+1, right)
		}

		if right < left {
			current[left+right] = ')'
			dfs(left, right+1)
		}
	}

	dfs(0, 0)
	return res
}
