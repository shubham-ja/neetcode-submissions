func maxSlidingWindow(nums []int, k int) []int {
    rightMax := make([]int, len(nums))
	leftMax := make([]int, len(nums))
	
	leftMax[0] = nums[0]
	rightMax[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i<len(nums); i++ {
		if i%k == 0 {
			leftMax[i] = nums[i]
		} else {
			leftMax[i] = max(leftMax[i-1],nums[i])
		}
	}
	for i := len(nums)-2; i >= 0 ; i-- {
		if i%k == 0 {
			rightMax[i] = nums[i]
		} else {
			rightMax[i] = max(rightMax[i+1],nums[i])
		}
	}
	var result []int 
	for i := k-1; i<len(nums); i++ {
		rightGroupMax := leftMax[i]
		leftGroupMax := rightMax[i-k+1]
		result = append(result,max(rightGroupMax, leftGroupMax))
	}
	return result
}
