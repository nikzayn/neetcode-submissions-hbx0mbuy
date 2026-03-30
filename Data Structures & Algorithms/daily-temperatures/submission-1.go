func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures))
	
	for i, j := range temperatures {
		for len(stack) > 0 && j > temperatures[stack[len(stack) - 1]] {
			stackIndex := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			res[stackIndex] = i - stackIndex
		}
		stack = append(stack, i)
	}

	return res
}
