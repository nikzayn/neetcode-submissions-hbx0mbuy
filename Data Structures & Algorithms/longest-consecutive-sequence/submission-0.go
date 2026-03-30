func longestConsecutive(nums []int) int {
	longest := 0

	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true
	}

	for val := range numSet {
		if !numSet[val - 1] {
			current := val
			length := 1

			for numSet[current + 1] {
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
