func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	elems := make(map[byte]int)
	for i := range(s) {
		a := s[i]
		b := t[i]

		elems[a]++
		elems[b]--

		if elems[a] == 0 {
			delete(elems, a)
		}

		if elems[b] == 0 {
			delete(elems, b)
		}
	}

	if len(elems) == 0 {
		return true
	}
	return false
}
