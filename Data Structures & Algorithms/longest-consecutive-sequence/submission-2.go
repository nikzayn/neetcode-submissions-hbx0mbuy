func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool)
	maxLen := 0

	for _, num := range nums {
		set[num] = true
	}

	for _, num := range nums {
		if !set[num - 1] {
			current := num
			length := 1

			for set[current + 1] {
				current++
				length++
			}

			maxLen = max(maxLen, length)
		}
	}

	return maxLen
}
