import (
	"slices"
)

func countReqTime(piles []int, s int) int {
	ans := 0
	for _, pile := range(piles) {
		ans += (pile + s - 1) / s
	}
	return ans
}

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, slices.Max(piles)

	for lo < hi {
		mid := lo + (hi - lo) / 2
		time := countReqTime(piles, mid)
		if time <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}
