func lemonadeChange(bills []int) bool {
	fives := 0
	tens := 0
	for _, bill := range bills {
		if bill == 5 {
			fives++
		} else if bill == 10 {
			if fives > 0 {
				tens++
				fives--
			} else {
				return false
			}
		} else {
			if tens > 0 && fives > 0 {
				tens--
				fives--
			} else if fives >= 3 {
				fives -= 3
			} else {
				return false
			}
		}
	}
	return true
}
