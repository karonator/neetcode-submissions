func isPalindrome(s string) bool {
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s) - 1 - i] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	result := [][]string{}

	var backtrack func(int, []string)
	backtrack = func(start int, cur []string) {
		if start == len(s) {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			result = append(result, tmp)
			return
		}

		for end := start; end < len(s); end++ {
			piece := s[start: end + 1]
			if !isPalindrome(piece) {
				continue
			}
			cur = append(cur, piece)
			backtrack(end + 1, cur)
			cur = cur[:len(cur) - 1]
		}
	}
	backtrack(0, []string{})
	return result
}
