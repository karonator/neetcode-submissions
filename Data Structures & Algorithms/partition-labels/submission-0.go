func partitionLabels(s string) []int {
    ends := make(map[byte]int)
	for i := range s {
		c := s[i]
		ends[c] = i
	}

	ans := make([]int, 0)
	start := 0
	end := 0
	for i := range s {
		c := s[i]
		end = max(end, ends[c])
		if i == end {
			ans = append(ans, end - start + 1)
			start = end + 1
		}
	}
	return ans

}
