func jump(nums []int) int {
    jumps, farthest, currEnd := 0, 0, 0

	for i := 0; i < len(nums) - 1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == currEnd {
			jumps++
			currEnd = farthest
		}
	}

	return jumps
}
