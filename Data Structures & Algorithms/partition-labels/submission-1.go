func partitionLabels(s string) []int {
	lastOccurence := make(map[byte]int)
	for i := range s {
		lastOccurence[s[i]] = i
	}
	ans := []int{}
	start := 0
	end := -1
	for i := range s {
		end = max(end, lastOccurence[s[i]])
		if i == end {
			ans = append(ans, end - start + 1)
			start = end + 1
		}
	}

	return ans
}
