func combinationSum2(candidates []int, target int) [][]int {
    var res [][]int

    sort.Ints(candidates)


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

        for i := start; i < len(candidates); i++ {
                if i > start && candidates[i] == candidates[i-1] {
                    continue
                }

                if candidates[i] > target {
                    continue
                }

                path = append(path, candidates[i])

                dfs(i+1, target - candidates[i], path)

                path = path[:len(path) - 1]
        }
    }

    dfs(0, target, []int{})
    return res
}
