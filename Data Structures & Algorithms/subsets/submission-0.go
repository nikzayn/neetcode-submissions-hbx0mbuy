func subsets(nums []int) [][]int {
    var res [][]int

    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        temp := make([]int, len(path))
        copy(temp, path)
        res = append(res, temp)

        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])

            dfs(i+1, path)

            path = path[:len(path) - 1]
        }
    }

    dfs(0, []int{})

    return res
}
