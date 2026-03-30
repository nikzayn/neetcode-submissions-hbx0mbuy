func countSubstrings(s string) int {

    count := 0

	for i := 0; i < len(s); i++ {
		count += expand(i, i, s)
		count += expand(i, i + 1, s)
	}

	return count
}

func expand(left, right int, s string) int {
		count := 0
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right++
		}
	return count
}
