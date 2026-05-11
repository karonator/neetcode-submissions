func isValid(c byte) bool {
	if c >= 'a' && c <= 'z' || c >='A' && c <= 'Z' || c >= '0' && c <= '9' {
		return true
	}
	return false
}

func lower(c byte) byte {
	if c >='A' && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if !isValid(s[i]) {
			i++
			continue
		}
		if !isValid(s[j]) {
			j--
			continue
		}

		if lower(s[i]) != lower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
