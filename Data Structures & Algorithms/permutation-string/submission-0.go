func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}

	if len(s1) > len(s2) {
		return false
	} 


	target, window := [26]int{}, [26]int{}

	for i := 0; i < len(s1); i++ {
		target[s1[i] - 'a']++
		window[s2[i] - 'a']++
	}

	fmt.Println(target, window)

	if target == window {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		window[s2[i] - 'a']++
		window[s2[i - len(s1)] - 'a']--

		if window == target {
			return true
		}
	}

	return false


}
