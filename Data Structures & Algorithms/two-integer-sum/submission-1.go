func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for id, num := range nums {
		rem := target - num
		if val, found := m[num]; found {
			return []int{val, id}
		}
		m[rem] = id
	}

	return []int{}
}
