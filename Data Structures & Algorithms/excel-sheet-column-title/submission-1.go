func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber > 0 {
		columnNumber--
		c := columnNumber % 26
		ans = string('A' + c) + ans
		columnNumber /= 26
	}
	return ans
}
