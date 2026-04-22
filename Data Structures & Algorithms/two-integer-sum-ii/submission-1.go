func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1

	for left < right {
		currSum := numbers[left] + numbers[right]

		if currSum > target {
			right--
		} else if currSum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{}
}
