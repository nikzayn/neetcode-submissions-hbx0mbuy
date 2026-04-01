func characterReplacement(s string, k int) int {
	start, maxLen, maxCount := 0, 0, 0

	freq := make(map[byte]int)

	for end := 0; end < len(s); end++ {
		freq[s[end]]++

		maxCount = max(maxCount, freq[s[end]])

		if (end - start + 1) - maxCount > k {
			freq[s[start]]--
			start++
		}

		maxLen = max(maxLen, end - start + 1)
	}

	return maxLen
}
