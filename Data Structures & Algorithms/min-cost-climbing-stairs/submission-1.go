func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	prev1, prev2 := cost[0], cost[1]

	for i := 2; i < n; i++ {
		curr := cost[i] + min(prev1, prev2)
		prev1 = prev2
		prev2 = curr
	}

	return min(prev1, prev2)
}
