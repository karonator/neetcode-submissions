import (
	"slices"
)

func shipWithinDays(weights []int, days int) int {
	lo, hi := slices.Max(weights), 0
	for _, weight := range(weights) {
		hi += weight
	}

	var is_ok func(int) bool
	is_ok = func(cap int) bool {
		left := 0
		nes_days := 0
		for _, weight := range(weights) {
			if left < weight {
				left = cap
				nes_days++
				if nes_days > days {
					return false
				}
			}
			left -= weight
		}
		return true
	}


	for lo < hi {
		mid := lo + (hi - lo) / 2
		if is_ok(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}