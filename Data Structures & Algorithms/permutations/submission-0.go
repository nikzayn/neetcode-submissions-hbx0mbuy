func permute(nums []int) [][]int {
    var res [][]int
    dfs(&res, nums, 0)
    return res
}

func dfs(res *[][]int, nums []int, idx int) {
    if idx == len(nums) {
        temp := append([]int{}, nums...)
        *res = append(*res, temp)
        return
    }

    for i := idx; i < len(nums); i++ {
        nums[idx], nums[i] = nums[i], nums[idx]
        dfs(res, nums, idx+1)
        nums[idx], nums[i] = nums[i], nums[idx]
    }
}
