func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	target := make(map[byte]int)
	for i := range s1 {
		target[s1[i]]++
	}

	window := make(map[byte]int)

	left := 0
	matches := 0
	needs := len(target)

	for right := 0; right < len(s2); right++ {
		c := s2[right]
		window[c]++

		if window[c] == target[c] {
			matches++
		} else if window[c]-1 == target[c] {
			matches--
		}

		if right-left+1 > len(s1) {
			l := s2[left]

			if window[l] == target[l] {
				matches--
			} else if window[l]-1 == target[l] {
				matches++
			}

			window[l]--
			left++
		}

		if matches == needs {
			return true
		}
	}

	return false
}