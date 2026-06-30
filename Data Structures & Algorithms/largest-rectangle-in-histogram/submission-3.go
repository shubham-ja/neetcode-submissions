func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	for i, h := range heights {
		for len(stack)>0 && heights[stack[len(stack)-1]] > h {
			topIndex := stack[len(stack)-1]
			topHeight := heights[topIndex]
			stack = stack[:len(stack)-1]
			width := i 
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := topHeight * width
			if area > maxArea {
				maxArea = area
			}
			
		}
		stack = append(stack,i)
		
	}
	for index, i := range stack {
		width := len(heights) - index
		
		if index > 0 {
			width = len(heights) - 1 - stack[index-1]
		}
		area := heights[i]*(width)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
