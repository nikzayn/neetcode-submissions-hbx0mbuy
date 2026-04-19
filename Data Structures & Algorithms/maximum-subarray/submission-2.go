func maxSubArray(nums []int) int {
    maxSum, maxCurr := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		maxCurr = max(nums[i], nums[i]+maxCurr)
		maxSum = max(maxSum, maxCurr)
	}

	return maxSum
}
