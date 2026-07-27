func isValid(s string) bool {
	stack := []rune{}
	closers := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range(s) {
		if _, found := closers[c]; !found {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack) - 1]
			stack = stack[0: len(stack) - 1]
			opener, _ := closers[c]
			if opener != last {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}

	return true
}
