func numDecodings(s string) int {
	decoder := make(map[string]struct{})
	for i := 1; i <= 26; i++ {
		decoder[strconv.Itoa(i)] = struct{}{}
	}
	
	decodings := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if _, found := decoder[s[i: i + 1]]; found {
			prev := 1
			if i > 0 {
				prev = decodings[i - 1]
			}
			decodings[i] += prev
		}
		if i > 0 {
			if _, found := decoder[s[i - 1: i + 1]]; found {
				pre_prev := 1
				if i > 1 {
					pre_prev = decodings[i - 2]
				}
				decodings[i] += pre_prev
			}
		}
	}

	return decodings[len(decodings) - 1]
}
