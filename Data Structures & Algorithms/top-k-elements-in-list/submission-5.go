func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))

	//frequency map
	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums)+1)

	//buckets by frequency
	for k, v := range freq {
		buckets[v] = append(buckets[v], k)
	}

	res := make([]int, 0, k)

	//traverse backwards -> take K
	for i := len(buckets) - 1; i >= 0 && k > 0; i-- {
		for _, v := range buckets[i] {
			res = append(res, v)
			k--
		}
	}

	return res
}
