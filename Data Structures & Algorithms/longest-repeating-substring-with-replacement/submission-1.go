func max (a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func characterReplacement(s string, k int) int {
	freqs := make([]int, 26)

	left := 0

	max_freq := 0
	max_win_len := 0

	rescan := func() {
		max_freq = 0
		for i := range(freqs) {
			max_freq = max(max_freq, freqs[i])
		}
	}

	for right := 0; right < len(s); right ++ {
		c := s[right] - 'A'
		freqs[c] ++

		if max_freq < freqs[c] {
			max_freq = freqs[c]
		}

		for right - left + 1 - max_freq > k {
			freqs[s[left] - 'A']--
			left++
			rescan()
		}

		max_win_len = max(max_win_len, right - left + 1)
	}
	return max_win_len
}
