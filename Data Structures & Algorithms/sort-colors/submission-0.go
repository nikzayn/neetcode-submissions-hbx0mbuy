func sortColors(nums []int) {
    sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}
