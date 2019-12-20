package climbingStairs

func climbStairs(n int) int {
	lastOneStepWays, lastTwoStepWays := 1, 0
	for i := 1; i < n; i++ {
		lastOneStepWays, lastTwoStepWays = lastOneStepWays+lastTwoStepWays, lastOneStepWays
	}
	return lastOneStepWays + lastTwoStepWays
}
