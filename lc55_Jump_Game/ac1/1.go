package ac1

// 动态规划
// 设dp[i]表示走到第i格时剩余的最大步数，则
// dp[i] = max(dp[i-1]-1, nums[i])
// 若遍历到最后一个之前dp[i]就变为0，则不能走完
func canJump(nums []int) bool {
	dp := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp <= 0 {
			return false
		}
		dp = max(dp-1, nums[i])
	}
	return true
}
