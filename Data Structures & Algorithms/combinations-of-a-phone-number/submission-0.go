func letterCombinations(digits string) []string {
	var res []string
	var path string

	if len(digits) == 0 {
		return []string{}
	}

	dfs(0, digits, path, &res)
	return res
}

func dfs(index int, digits, path string, res *[]string) {
	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	if index == len(digits) {
		*res = append(*res, path)
		return
	}

	letters := phone[digits[index]]

	for i := 0; i < len(letters); i++ {
		dfs(index+1, digits, path+string(letters[i]), res)
	}
}