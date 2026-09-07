func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := -1
	end := -1
	ans := make([][]int, 0)

	for _, interval := range intervals {
		if start < 0 {
			start = interval[0]
			end = interval[1]
		} else {
			if interval[0] <= end {
				end = max(end, interval[1])
			} else {
				ans = append(ans, []int{start, end})
				start = interval[0]
				end = interval[1]
			}
		}
	}
	if start >= 0 {
		ans = append(ans, []int{start, end})
	}	
	return ans
}
