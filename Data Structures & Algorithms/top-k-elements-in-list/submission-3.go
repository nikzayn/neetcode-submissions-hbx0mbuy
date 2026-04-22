func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))

	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums)+1)

	for key, val := range freq{
		buckets[val] = append(buckets[val], key)
	}

	res := make([]int, 0, k)

	for i := len(buckets) - 1; i > 0; i-- {
		for _, v := range buckets[i] {
			if k > 0 {
				res = append(res, v)
				k--
			}
		}
	}

	return res
}
