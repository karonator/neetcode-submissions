func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	window := make(map[byte]int)

	ans := 0
	left := 0
	for right := 0; right < len(s); right ++ {
		window[s[right]] ++

		for window[s[right]] > 1 {
			window[s[left]] --
			left ++
		}

		ans = max(ans, right - left + 1)
	}

	return ans
}
