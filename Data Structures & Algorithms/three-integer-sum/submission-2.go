func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)

	n := len(nums)

	if n == 0 {
		return res
	}

	for i := 0; i <= n - 2; i++ {
		left, right := i + 1, n - 1

		if i > 0 && nums[i - 1] == nums[i] {
			continue
		}

		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[left], nums[right], nums[i]})

				for left < right && nums[left] == nums[left + 1] {
					left++
				}

				for left > right && nums[right - 1] == nums[left] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return res
}
