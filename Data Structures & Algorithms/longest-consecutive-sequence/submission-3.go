func longestConsecutive(nums []int) int {
	freq := make(map[int]bool)

	for _, num := range nums {
		freq[num] = true
	}

	maxLen := 0

	for _, num := range nums {
		if !freq[num - 1] {
			current := num
			length := 1

			for freq[current + 1] {
				current++
				length++
			}

			maxLen = max(maxLen, length)
		}
	}

	return maxLen
}
