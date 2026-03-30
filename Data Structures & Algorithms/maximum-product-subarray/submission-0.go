func maxProduct(nums []int) int {
	res, currMax, currMin := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if num < 0 {
			currMax, currMin = currMin, currMax
		}
		currMax = max(num, num * currMax)
		currMin = min(num, num * currMin)
		res = max(res, currMax)
	}

	return res
}
