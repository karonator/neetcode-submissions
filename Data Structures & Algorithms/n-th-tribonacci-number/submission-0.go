func tribonacci(n int) int {
	data := [3]int{0, 1, 1}
	for i := 3; i <= n; i++ {
		ans := data[0] + data[1] + data[2]
		data[0] = data[1]
		data[1] = data[2]
		data[2] = ans
	}

	if n <= 1 {
		return data[n]
	}
	return data[2]
}
