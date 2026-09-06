func isAlienSorted(words []string, order string) bool {
	alphabet := make(map[byte]int, len(order))
	for i := range(order) {
		alphabet[order[i]] = i
	}

	sorted := func(s1 string, s2 string) bool {
		i := 0
		for i < min(len(s1), len(s2)) {
			if alphabet[s1[i]] > alphabet[s2[i]] {
				return false
			} else if alphabet[s1[i]] < alphabet[s2[i]] {
				return true
			}
			i++
		}
		if len(s1) > len(s2) {
			return false
		}
		return true
	}

	for i := 0; i < len(words) - 1; i++ {
		if !sorted(words[i], words[i + 1]) {
			return false
		}
	}
	return true
}
