func lengthOfLongestSubstring(s string) int {
	window_chars := make(map[byte]int)
	ans := 0
	left := 0

	for right := 0; right < len(s); right++ {
		c := s[right]

		window_chars[c]++

		for left <= right && window_chars[c] > 1 {
			window_chars[s[left]]--
			left++
		}

		tmp_ans := right - left + 1
		if tmp_ans > ans {
			ans = tmp_ans
		}
	}

	return ans
}
