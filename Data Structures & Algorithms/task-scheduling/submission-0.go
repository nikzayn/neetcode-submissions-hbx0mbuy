func leastInterval(tasks []byte, n int) int {
    freqMap := make([]int, 26)

    for _, task := range tasks {
        freqMap[task - 'A']++
    }

    maxFreq := 0
    for _, f := range freqMap {
        if f > maxFreq {
            maxFreq = f
        }
    }

    maxCount := 0
    for _, f := range freqMap {
        if f == maxFreq {
            maxCount++
        }
    }

    groupInterval := maxFreq - 1
    sortLength := n + 1
    emptySlots := groupInterval * sortLength + maxCount

    return max(len(tasks), emptySlots)
}
