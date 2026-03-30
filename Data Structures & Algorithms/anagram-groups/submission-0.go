func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string, len(strs))

	for _, str := range strs {
		sortedWord := sortedWord(str)
		anagrams[sortedWord] = append(anagrams[sortedWord], str)
	}

	res := [][]string{}
	for _, val := range anagrams {
		res = append(res, val)
	}

	return res
}

func sortedWord(str string) string{
	words := []rune(str)

	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	return string(words)
}