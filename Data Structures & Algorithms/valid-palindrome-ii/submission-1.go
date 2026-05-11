func core(s string, allowFail bool) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			if !allowFail {
				return false
			} else {
				return core(s[i + 1: j + 1], false) || core(s[i: j], false) 
			}
		}
		i++
		j--
	}
	return true
}

func validPalindrome(s string) bool {
	return core(s, true)
}