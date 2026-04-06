func majorityElement(nums []int) []int {
	res := []int{}

	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	for key, val := range freq {
		if val > len(nums) / 3 {
			res = append(res, key)
		}
	}

	return res
}
