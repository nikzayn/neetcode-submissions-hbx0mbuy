func lengthOfLongestSubstring(s string) int {
    start, maxLen := 0, 0
    chars := make(map[byte]int)

    for end := 0; end < len(s); end++ {
        chars[s[end]]++

        for chars[s[end]] > 1 {
            chars[s[start]]--
            start++
        }

        maxLen = max(maxLen, end - start + 1)
    }
    
    return maxLen
}
