func partition(s string) [][]string {
	var res [][]string
	var path []string

	dfs(0, path, s, &res)
	return res
}
func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func dfs(start int, path []string, s string, res *[][]string) {
	if start == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			dfs(i+1, path, s, res)
			path = path[:len(path) - 1]
		}
	}
}
