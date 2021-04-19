package ac1

// 动态规划
// 设dp[i]表示以nums[i]结尾的最长上升子序列的长度，则状态转移方程为
// dp[i] = max(dp[j]) + 1, 0 <= j < i, nums[j] < nums[i]
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	ans := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] + 1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > ans {
					ans = dp[i]
				}
			}
		}
	}
	return ans
}
