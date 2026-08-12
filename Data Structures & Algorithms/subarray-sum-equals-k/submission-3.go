func subarraySum(nums []int, k int) int {
	freq := make(map[int]int)

	count, prefixSum := 0, 0
	freq[0] = 1

	for _, num := range nums {
		prefixSum += num

		if val, found := freq[prefixSum - k]; found {
			count += val
		}

		freq[prefixSum]++
	}

	return count
}
