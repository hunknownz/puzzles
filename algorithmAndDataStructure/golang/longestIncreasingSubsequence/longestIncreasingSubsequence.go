package longestIncreasingSubsequence

func lower_bound(nums []int, value int) (x int) {
	y := len(nums)
	for x < y {
		m := x + (y-x)/2
		if nums[m] >= value {
			y = m
		} else {
			x = m + 1
		}
	}
	return
}

func lengthOfLIS(nums []int) int {
	numsLen := len(nums)
	lowEndings := make([]int, 0, numsLen)
	for _, num := range nums {
		i := lower_bound(lowEndings, num)
		lowEndingsLen := len(lowEndings)
		if i >= lowEndingsLen {
			lowEndings = append(lowEndings, num)
		} else {
			lowEndings[i] = num
		}
	}
	return len(lowEndings)
}
