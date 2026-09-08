func numDecodings(s string) int {
    data := make([]int, len(s))
	if s[0] == '0' {
		return 0
	}
	data[0] = 1

	for i := 1; i < len(s); i++ {
		cur := s[i] - '0'
		pre := s[i - 1] - '0'

		if cur > 0 {
			data[i] = data[i - 1]
		}
		if pre == 1 || pre == 2 && cur <= 6 {
			if i > 1 {
				data[i] += data[i - 2]
			} else {
				data[i] += 1
			}
		}
		if data[i] == 0 {
			return 0
		}
	}
	return data[len(data) - 1]
}
