func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string, len(strs))

	for _, str := range strs {
		sortWord := sortedWord(str)
		anagrams[sortWord] = append(anagrams[sortWord], str)
	}

	result := [][]string{}
	for _, anagram := range anagrams {
		result = append(result, anagram)
	}

	return result
}

func sortedWord(str string) string {
	words := []rune(str)
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	return string(words)
}
