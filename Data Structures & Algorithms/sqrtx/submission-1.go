func sqrt(a int) int {
	return a * a
}

func mySqrt(x int) int {
	start := 1
	end := x

	for start <= end {
		mid := start + (end - start) / 2

		if sqrt(mid) == x {
			return mid
		}

		if sqrt(mid) > x {
			end = mid - 1 
		} else {
			start = mid + 1
		}

	}
	return end
}
