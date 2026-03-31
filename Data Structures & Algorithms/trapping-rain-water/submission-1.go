func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	left, right := 0, n - 1
	leftMax, rightMax, ans := 0, 0, 0

	for left <= right {
		if height[left] < height[right] {
			if leftMax > height[left] {
				ans = ans + (leftMax - height[left])
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if rightMax > height[right] {
				ans = ans + (rightMax - height[right])
			} else {
				rightMax = height[right]
			}
			right--
		}
	}
	return ans
}
