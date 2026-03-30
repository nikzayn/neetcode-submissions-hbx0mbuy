func getConcatenation(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < 2; i++ {
		for _, num := range nums {
			res = append(res, num)
		}
	}

	return res
}
