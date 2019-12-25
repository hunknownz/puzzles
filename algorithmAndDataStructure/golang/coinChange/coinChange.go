package coinChange

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	coinsLen := len(coins)
	for i := 0; i <= amount; i++ {
		for j := 0; j <= coinsLen; j++ {
			if i < coins[j] {
				continue
			}
			tmp := dp[i-coins[j]]
			if tmp == -1 {
				continue
			}
			if dp[i] == -1 || dp[i] < tmp+1 {
				dp[i] = tmp + 1
			}
		}
	}
	return dp[amount]
}
