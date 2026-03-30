func rob(nums []int) int {
	n := len(nums)
    prev1, prev2 := 0, 0
	if n == 1 {
		return nums[0]
	}

	for i := 0; i < len(nums); i++ {
		// For each house, get max robbery amount
		curr := max(prev1, nums[i]+prev2)
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}
