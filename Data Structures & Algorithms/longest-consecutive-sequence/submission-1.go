func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	freq := make(map[int]bool, len(nums))

	for _, num := range nums {
		freq[num] = true
	}

	longest := 0

	for val := range freq {
		if !freq[val - 1] {
			current := val
			length := 1

			for freq[current + 1] {
				current++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
