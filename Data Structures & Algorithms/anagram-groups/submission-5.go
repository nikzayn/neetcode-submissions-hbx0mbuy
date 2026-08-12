func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[string][]string, len(strs))

    for _, str := range strs {
        sortedWord := sortWord(str)
        anagrams[sortedWord] = append(anagrams[sortedWord], str)
    }

    res := [][]string{}

    for _, anagram := range anagrams {
        res = append(res, anagram)
    }

    return res
}

func sortWord(str string) string {
    words := []rune(str)
    sort.Slice(words, func(i, j int) bool {
        return words[i] < words[j]
    })

    return string(words)
}
