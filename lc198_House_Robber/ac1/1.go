package ac1

// 动态规划
// dp[i] 表示下标范围 [1,i] 内能偷的最大价值，则
// dp[i] = max(nums[i-2]+nums[i], nums[i-1)
// dp[len(nums)-1] 即为最终结果
func rob(nums []int) int {
	// x表示i的前两个数，y表示i的前一个数
	x, y, z := 0, 0, 0
	for i := range nums {
		z = max(x + nums[i], y)
		x, y = y, z
	}
	return z
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
