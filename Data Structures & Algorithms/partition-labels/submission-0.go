func partitionLabels(s string) []int {
    last := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		last[s[i]] = i
	}

	res := []int{}
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		end = max(end, last[s[i]])
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}

	return res

}
