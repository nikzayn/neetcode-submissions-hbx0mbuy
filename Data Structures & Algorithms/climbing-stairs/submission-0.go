func climbStairs(n int) int {
    if n <= 2 {
		return n
	}

	prev2, prev1 := 1, 2

	//ways to achieve n ways for climbing stairs
	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}
