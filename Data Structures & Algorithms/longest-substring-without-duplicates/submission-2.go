func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	start, maxLen := 0, 0

	freq := make(map[byte]int)

	for end := 0; end < len(s); end++ {
		freq[s[end]]++

		for freq[s[end]] > 1 {
			freq[s[start]]--
			start++
		}

		maxLen = max(maxLen, end - start + 1)
	}

	return maxLen
}
