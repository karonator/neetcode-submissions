type stack []rune

func (s stack) Push(c rune) stack {
	return append(s, c)
}

func (s stack) Pop() (stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func isValid(s string) bool {
	pair := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}

	openers := map[rune]struct{}{
		'[': {},
		'(': {},
		'{': {},
	}

	stk := make(stack, 0)

	for _, c := range s {
		if _, isOpener := openers[c]; isOpener {
			stk = stk.Push(c)
		} else {
			if len(stk) == 0 {
				return false
			}

			var opener rune
			stk, opener = stk.Pop()

			if opener != pair[c] {
				return false
			}
		}
	}

	return len(stk) == 0
}