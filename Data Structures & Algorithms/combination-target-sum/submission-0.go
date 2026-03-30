func combinationSum(nums []int, target int) [][]int {
    var res [][]int
    
    var dfs func(start, target int, path []int)
    dfs = func(start, target int, path []int) {
            if target == 0 {
                temp := make([]int, len(path))
                copy(temp, path)
                res = append(res, temp)
                return
            }

            if target < 0 {
                return
            }

            for i := start; i < len(nums); i++ {
                path = append(path, nums[i])

                dfs(i, target - nums[i], path)

                path = path[:len(path) - 1]
            }
    }

    dfs(0, target, []int{})
    return res
}
