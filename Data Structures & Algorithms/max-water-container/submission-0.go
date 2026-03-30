func maxArea(heights []int) int {
	left, right := 0, len(heights) - 1

	maxArea := 0

	for left < right {
		currArea := min(heights[left], heights[right]) * (right - left)
		
		if heights[left] < heights[right] {
			maxArea = max(maxArea, currArea)
			left++
		} else {
			maxArea = max(maxArea, currArea)
			right--
		}
	}

	return maxArea
}
