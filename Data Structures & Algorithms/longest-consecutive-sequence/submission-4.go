import (
    "slices"
)
func longestConsecutive(nums []int) int {
    if len(nums) < 1 {
        return 0
    }
    slices.Sort(nums)
    maxLength := 1
    currentLength := 1
    for i := 1; i < len(nums); i++ {
        elem := nums[i];
        diff := elem-nums[i-1];
        if diff == 0 {
            continue
        } else if diff == 1 {
            currentLength += 1
            if maxLength < currentLength {
                maxLength = currentLength
            }
        } else {
            currentLength = 1
        }
    }
    return maxLength
}
