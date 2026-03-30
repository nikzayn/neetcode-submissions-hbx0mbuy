func characterReplacement(s string, k int) int {
	start, maxLen, maxCount := 0, 0, 0

	chars := make(map[byte]int)

	for end := 0; end < len(s); end++ {
		chars[s[end]]++

		maxCount = max(maxCount, chars[s[end]])

		if (end - start + 1) - maxCount > k {
			chars[s[start]]--
			start++
		}

		maxLen = max(maxLen, end - start + 1)
	}

	return maxLen
}
