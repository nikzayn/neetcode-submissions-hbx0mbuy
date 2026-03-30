func threeSum(nums []int) [][]int {

	result := [][]int{}

	if len(nums) == 0 {
		return result
	}

	sort.Ints(nums)

	n := len(nums)

	for i := 0; i < n - 2; i++ {
		low, high := i + 1, n - 1

		if i > 0 && nums[i - 1] == nums[i] {
			continue
		}

		for low < high {

			sum := nums[low] + nums[high] + nums[i]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[low], nums[high]})


				for low < high && nums[low] == nums[low + 1] {
					low++
				}

				for low > high && nums[high - 1] == nums[low] {
					high--
				}

				low++
				high--
			} else if sum < 0 {
				low++
			} else {
				high--
			}
		}
	}
	return result
}
