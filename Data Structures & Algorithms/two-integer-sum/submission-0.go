func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for idx, num := range nums {
		rem := target - num
		if val, exists := m[num]; exists {
			return []int{val, idx}
		}
		m[rem] = idx
	}

	return []int{}
}
