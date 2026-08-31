func longestCommonPrefix(strs []string) string {
    for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for _, s := range(strs) {
			if len(s) <= i {
				return strs[0][:i]
			}
			if s[i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
