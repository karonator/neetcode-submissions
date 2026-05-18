func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCostClimbingStairs(cost []int) int {
    path := make([]int, len(cost))
	for i := range cost {
		prev := 0
		if i > 0 {
			prev = path[i - 1]
		}
		pre_prev := 0
		if i > 1 {
			pre_prev = path[i - 2]
		}
		path[i] = cost[i] + min(prev, pre_prev)
	}

	return min(path[len(path) - 1], path[len(path) - 2])
}
