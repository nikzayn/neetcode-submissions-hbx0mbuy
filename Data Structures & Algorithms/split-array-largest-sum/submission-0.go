func splitArray(nums []int, k int) int {
	l, r := maxInArray(nums), sumArray(nums)

	for l < r {
		mid := (l + r) / 2
		if canSplit(mid, k, nums) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func canSplit(maxSum, k int, nums []int) bool {
	splits, currentSum := 1, 0

	for _, num := range nums {
		if currentSum + num > maxSum {
			splits++
			currentSum = num
		} else {
			currentSum += num
		}
	}

	return splits <= k
}

func maxInArray(nums []int) int {
	maxVal := nums[0]

	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	return maxVal
}

func sumArray(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}
