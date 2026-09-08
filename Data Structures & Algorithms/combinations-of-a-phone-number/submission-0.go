func letterCombinations(digits string) []string {
	digiMap := map[byte][]byte {
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	if len(digits) == 0 {
		return []string{}
	}

	ans := make([]string, 0)
	cur := make([]byte, 0)

	var backtrack func(int, []byte)
	backtrack = func(i int, cur []byte) {
		if i == len(digits) {
			ans = append(ans, string(cur))
			return
		}

		for _, c := range digiMap[digits[i]] {
			cur = append(cur, c)
			backtrack(i + 1, cur)
			cur = cur[:len(cur) - 1]
		}
	}
	backtrack(0, cur)
	return ans
}
