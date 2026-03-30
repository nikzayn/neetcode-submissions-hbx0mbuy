func largestRectangleArea(heights []int) int {
    stack := []int{}
    maxArea := 0
    heights = append(heights, 0)

    for i := 0; i < len(heights); i++ {
        for len(stack) > 0 && heights[i] < heights[stack[len(stack) - 1]] {
            h := heights[stack[len(stack) - 1]]
            stack = stack[:len(stack) - 1]
            var w int
            if len(stack) == 0 {
                w = i
            } else {
                w = i - stack[len(stack) - 1] - 1
            }
            area := h * w
            if area > maxArea {
                maxArea = area
            }
        }
        stack = append(stack, i)
    }
    return maxArea
}
