func insert(intervals [][]int, newInterval []int) [][]int {
    intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)
	start := intervals[0][0]
	end := intervals[0][1]
	
	for i := 1; i < len(intervals); i++ {
		if end >= intervals[i][0] {
			end = max(end, intervals[i][1])
		} else {
			ans = append(ans, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}
