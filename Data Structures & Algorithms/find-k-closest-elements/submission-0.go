func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func findClosestElements(arr []int, k int, x int) []int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + (hi - lo) / 2
		if x <= arr[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	right := lo
	left := lo - 1
	for i := 0; i < k; i++ {
		if left < 0 {
			right++
		} else if right >= len(arr) {
			left--
		} else {
			if abs(arr[left] - x) <= abs(arr[right] - x) {
				left--
			} else {
				right++
			}
		}
	}
	return arr[left+1 : right]
}
