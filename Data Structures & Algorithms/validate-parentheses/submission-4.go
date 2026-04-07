func isValid(s string) bool {
    if len(s) == 0 || len(s) % 2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, char := range s {
		if _, found := pairs[char]; found {
			stack = append(stack, char)
		} else if len(stack) == 0 || pairs[stack[len(stack) - 1]] != char {
			return false
		} else {
			stack = stack[:len(stack) - 1]
		}
	}

	return len(stack) == 0
}
