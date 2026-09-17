func isValid(s string) bool {
	openers := map[byte]struct{}{
		'(': struct{}{},
		'{': struct{}{},
		'[': struct{}{},
	}
	closers := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]byte, 0)
	for i := range s {
		c := s[i]
		if _, found := openers[c]; found {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack) - 1]
			if last != closers[c] {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
