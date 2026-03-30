func minWindow(s string, t string) string {
    if t == "" {
		return ""
	}

	countT := make(map[rune]int)

	// filling the map
	for _, c := range t {
		countT[c]++
	}

	have, need := 0, len(countT)
	res := []int{-1, -1}
	resLen := math.MaxInt32
	l := 0

	window := make(map[rune]int)

	for r := 0; r < len(s); r++ {
		c := rune(s[r])
		window[c]++

		//Check if the length is available in left
		if countT[c] > 0 && countT[c] == window[c] {
			have++
		}

		for have == need {
			if (r - l + 1) < resLen {
				res = []int{l, r}
				resLen = r - l + 1
			}
			window[rune(s[l])]--
			if countT[rune(s[l])] > 0 && window[rune(s[l])] < countT[rune(s[l])] {
				have--
			}
			l++
		}
	}

	if res[0] == -1 {
		return ""
	}

	return s[res[0]:res[1]+1]
}
