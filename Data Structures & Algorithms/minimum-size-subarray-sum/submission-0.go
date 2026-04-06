func minSubArrayLen(target int, nums []int) int {
	start, total, res := 0, 0, len(nums) + 1

	for end := 0; end < len(nums); end++ {
		total += nums[end]

		for total >= target {
			if end - start + 1 < res {
				res = end - start + 1
			}
			total -= nums[start]
			start++
		}
	}

	if res == len(nums) + 1 {
		return 0
	}

	return res
}
