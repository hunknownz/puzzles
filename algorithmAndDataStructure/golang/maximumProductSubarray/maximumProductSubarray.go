package maximumProductSubarray

func min(x int, others ...int) (minValue int) {
	minValue = x
	for _, other := range others {
		if minValue > other {
			minValue = other
		}
	}
	return
}

func max(x int, others ...int) (maxValue int) {
	maxValue = x
	for _, other := range others {
		if maxValue < other {
			maxValue = other
		}
	}
	return
}

func maxProduct(nums []int) int {
	prevMin, prevMax := nums[0], nums[0]
	curMin, curMax := nums[0], nums[0]
	maxValue := nums[0]

	numsLen := len(nums)
	for i := 1; i < numsLen; i++ {
		curMin = min(prevMin*nums[i], prevMax*nums[i], nums[i])
		curMax = max(prevMin*nums[i], prevMax*nums[i], nums[i])
		maxValue = max(maxValue, curMax)
		prevMin, prevMax = curMin, curMax
	}
	return maxValue
}
