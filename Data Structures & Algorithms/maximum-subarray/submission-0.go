func maxSubArray(nums []int) int {
    maxSum, currMax := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currMax = max(nums[i], nums[i] + currMax)
		maxSum = max(maxSum, currMax)
	}

	return maxSum
}
