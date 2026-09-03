func checkInclusion(s1 string, s2 string) bool {
	freq := make([]int, 26)
	rules := 0

	for i := range(s1) {
		idx := s1[i] - 'a'
		if freq[idx] == 0 {
			rules++
		}
		freq[idx]++
	}

	window := make([]int, 26)
	
	left := 0
	for right := range s2 {
		if right - left == len(s1) {
			left_idx := s2[left] - 'a'
			if freq[left_idx] == window[left_idx] {
				rules++
			}
			window[left_idx]--
			if freq[left_idx] == window[left_idx] {
				rules--
			}
			left++
		}

		right_idx := s2[right] - 'a'
		if freq[right_idx] == window[right_idx] {
			rules++
		}
		window[right_idx]++
		if freq[right_idx] == window[right_idx] {
			rules--
		}

		if rules == 0 {
			return true
		}
	}
	return false
}
