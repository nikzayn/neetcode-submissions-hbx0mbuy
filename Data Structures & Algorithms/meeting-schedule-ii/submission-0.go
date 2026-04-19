/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	var time [][]int

	for _, i := range intervals {
		time = append(time, []int{i.start, 1})
		time = append(time, []int{i.end, -1})
	}

	sort.Slice(time, func(i, j int) bool {
		if time[i][0] == time[j][0] {
			return time[i][1] < time[j][1]
		}
		return time[i][0] < time[j][0]
	})

	res, count := 0, 0
	for _, t := range time {
		count += t[1]
		if count > res {
			res = count
		}
	}

	return res
}
