func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		diaEven := min(i, len(s) - i - 1)
		for j := 0; j <= diaEven; j++ {
			if s[i - j] != s[i + j] {
				break
			} else if 2 * j + 1 > len(ans) {
				ans = s[i - j: i + j + 1]
			}
		}

		diaOdd := min(i, len(s) - i - 2)
		for j := 0; j <= diaOdd; j++ {
			if s[i - j] != s[i + 1 + j] {
				break
			}
			if 2 * j + 2 > len(ans) {
				ans = s[i - j: i + j + 2]
			}
		}
	}
	return ans
}
