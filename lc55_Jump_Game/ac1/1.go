package ac1

// 动态规划
// 设dp[i]表示走到第i格时剩余的最大步数，则
// dp[i] = max(dp[i-1]-1, nums[i])
// 若遍历到最后一个之前dp[i]就变为0，则不能走完
func canJump(nums []int) bool {
	if len(nums) != 1 && nums[0] <= 0 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1]-1, nums[i])
		if i != len(nums) - 1 && nums[i] <= 0 {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
