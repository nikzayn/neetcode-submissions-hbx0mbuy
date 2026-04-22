func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string, len(strs))

	for _, str := range strs {
		sortedWord := sortWord(str)
		anagrams[sortedWord] = append(anagrams[sortedWord], str)
	}

	result := [][]string{}

	for _, anagram := range anagrams {
		result = append(result, anagram)
	}

	return result
}

func sortWord(s string) string {
	words := []rune(s)

	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	return string(words)
}