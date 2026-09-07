/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))

	for i, interval := range intervals {
		start[i] = interval.start
		end[i] = interval.end
	}

	sort.Ints(start)
	sort.Ints(end)

	rooms := 0
	maxRooms := 0

	i := 0
	j := 0

	for i < len(start) && j < len(start) {
		if start[i] < end[j] {
			rooms++
			i++
		} else if start[i] == end[j] {
			i++
			j++
		} else {
			rooms--
			j++
		}
		maxRooms = max(rooms, maxRooms)
	}
	return maxRooms
}
