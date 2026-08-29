func clean(s string) string {
	var buffer bytes.Buffer

	for i := range(len(s)) {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			buffer.WriteByte(s[i])
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			buffer.WriteByte('a' + s[i] - 'A')
		}
	}

	return buffer.String()
}


func isPalindrome(s string) bool {
	s = clean(s)
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s) - 1 - i] {
			return false
		}
	}
	return true
}
