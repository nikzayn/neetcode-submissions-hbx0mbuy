func productExceptSelf(nums []int) []int {
	n := len(nums)

	ans := make([]int, n)

	ans[0] = 1

	//Left to right
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i - 1] * nums[i - 1]
	}

	suffix := 1

	//Right to left
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suffix
		suffix *= nums[i]
	}

	return ans
}
