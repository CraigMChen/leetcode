package ac1

// 动态规划
// 设 dp[i] 表示组成总和为i的数的组合数量，则
// dp[i] = ∑j dp[i-nums[j]]
// 与lc518类似，由于本题中不同序列视为不同的组合，所以要先遍历i
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target + 1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
