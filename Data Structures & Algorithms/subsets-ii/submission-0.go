func subsetsWithDup(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)

	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			
			path = append(path, nums[i])

			dfs(i+1, path)

			path = path[:len(path) - 1]
		}
	}

	dfs(0, []int{})

	return res
}