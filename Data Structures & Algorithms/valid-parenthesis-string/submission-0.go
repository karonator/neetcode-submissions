func checkValidString(s string) bool {
	// мин кол-во открывающихся скобок
	lo := 0
	// макс кол-во открывающихся скобок
	hi := 0

	for i := range s {
		if s[i] == '(' {
			lo++
			hi++
		} else if s[i] == ')' {
			lo--
			hi--
		} else {
			lo--
			hi++
		}
		lo = max(lo, 0)
		if hi < 0 {
			return false
		}
	}
	return lo == 0
}
