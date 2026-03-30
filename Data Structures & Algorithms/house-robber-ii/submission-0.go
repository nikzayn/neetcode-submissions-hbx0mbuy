func rob(nums []int) int {
    n := len(nums)
	if n == 1 {
		return nums[0]
	}

	return max(robLinear(nums[:n-1]), robLinear(nums[1:]))
}

func robLinear(nums []int) int {
	prev2, prev1 := 0, 0

	for _, num := range nums {
		curr := max(prev1, num + prev2)
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}
