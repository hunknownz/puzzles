package maximumSubarray

import (
	"math"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	sum, maxValue := 0, math.MinInt32

	for _, num := range nums {
		sum += num
		maxValue = max(maxValue, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxValue
}
