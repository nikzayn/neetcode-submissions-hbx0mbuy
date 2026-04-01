func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 {
		return ""
	}

	countT := make(map[rune]int)

	for _, c := range t {
		countT[c]++
	}

	have, need := 0, len(countT)
	res := []int{-1, -1}
	resLen := math.MaxInt32
	l := 0

	window := make(map[rune]int)

	for r := 0; r < len(s); r++ {
		end := rune(s[r])
		window[end]++

		//towards right
		if countT[end] > 0 && window[end] == countT[end] {
			have++
		}

		for have == need {
			if (r - l + 1) < resLen {
				res = []int{l, r}
				resLen = r - l + 1
			}

			//towards left
			start := rune(s[l])
			window[start]--

			if countT[start] > 0 && window[start] < countT[start] {
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
