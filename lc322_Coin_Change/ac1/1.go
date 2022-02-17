package ac1

// 动态规划
// 设 dp[i] 表示总额为i的硬币数量最小值
// dp[i] = min(dp[i-coins[j]) + 1 (i-coins[j]>=0)
// dp[0] = 0
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := 1<<31
		for j := 0; j < len(coins); j++ {
			if i - coins[j] >= 0 && dp[i - coins[j]] < min && dp[i - coins[j]] != -1 {
				min = dp[i - coins[j]]
			}
		}
		if min == 1<<31 {
			dp[i] = -1
		} else {
			dp[i] = min + 1
		}
	}
	return dp[amount]
}
