func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := make([]int, 0)

	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	for k, v := range freq {
		if v > len(nums) / 3 {
			res = append(res, k)
		}
	}

	return res
}
