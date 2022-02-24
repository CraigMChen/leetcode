package ac1

// 动态规划
// 设 dp[i] 表示组成硬币i的组合数，则
// dp[i] = ∑j dp[i-coin[j]]
// 但是，先遍历i会导致重复计算，例如
// coins = [1, 2, 3], amount = 3 时
// 3 = 1 + 1 + 1
// 3 = 2 + 1
// 3 = 1 + 2
// 所以要先遍历j，即先遍历硬币的面额
func change(amount int, coins []int) int {
	dp := make([]int, amount + 1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j - coins[i]]
		}
	}
	return dp[amount]
}