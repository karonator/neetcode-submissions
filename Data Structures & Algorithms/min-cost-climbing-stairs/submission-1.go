func minCostClimbingStairs(cost []int) int {
    stair := make([]int, len(cost) + 1)

	stair[0], stair[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		stair[i] = min(stair[i - 1] + cost[i - 1], stair[i - 2] + cost[i - 2])
	}

	return stair[len(cost)]
}
