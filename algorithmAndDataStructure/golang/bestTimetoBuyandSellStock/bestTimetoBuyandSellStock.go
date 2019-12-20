package bestTimetoBuyandSellStock

import (
	"math"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxProfit(prices []int) int {
	minBuy, maxProfit := math.MaxInt32, 0
	for _, price := range prices {
		if price > minBuy {
			maxProfit = max(maxProfit, price-minBuy)
		}
		minBuy = min(minBuy, price)
	}
	return maxProfit
}
