func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    i := 0
    var res [][]int

    //we add all the intervals that add before new interval
    for i < n && intervals[i][1] < newInterval[0] {
        res = append(res, intervals[i])
        i++
    }

    //now we check the newInterval where it end should be greater than interval start
    for i < n && newInterval[1] >= intervals[i][0] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    res = append(res, newInterval)

    for i < n {
        res = append(res, intervals[i])
        i++
    }

    return res
}
