func minWindow(s string, t string) string {
    target := make(map[byte]int)
	for i := range t {
		target[t[i]]++
	}
	
	window := make(map[byte]int)
	start := 0
	done := 0
	ans := ""
	for i := range s {
		c := s[i]
		if window[c] == target[c] - 1 {
			done++
		}
		window[c]++
		for done == len(target) {
			if len(ans) > len(s[start: i + 1]) || ans == "" {
				ans = s[start: i + 1]
			}
			c := s[start]
			if window[c] == target[c] {
				done--
			}
			window[c]--
			start++
		}
	}
	return ans
}
