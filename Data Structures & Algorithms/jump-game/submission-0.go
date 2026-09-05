func canJump(nums []int) bool {
    reachable := 0
	for i, num := range nums {
		if i > reachable {
			return false
		}
		reachable = max(reachable, num + i)
	}
	return true
}
