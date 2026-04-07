func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string, len(strs))

	for _, str := range strs {
		sortedWord := sortWord(str)
		anagrams[sortedWord] = append(anagrams[sortedWord], str)
	}

	results := [][]string{}
	for _, anagram := range anagrams {
		results = append(results, anagram)
	}

	return results
}

func sortWord(s string) string {
	word := []rune(s)
	sort.Slice(word, func(i, j int) bool {
		return word[i] < word[j]
	})
	return string(word)
}