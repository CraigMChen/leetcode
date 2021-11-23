package ac1

// 动态规划
// 设 dp1[i] 表示偷第0间房子，即偷取范围为[0,n-1)，偷到第i间房子获得的最大价值
// dp0[i] 表示不偷第0间房子，即偷取范围为[1,n)，偷到第i间房子获得的最大价值，则
// dp1[0] = nums[0], dp1[1] = nums[0]
// dp0[0] = 0, dp0[1] = nums[1]
// dp[i] = max(dp[i-2]+nums[i], nums[i-1]
// 则 max(dp0[n-1], dp1[n-1]) 即为最终结果
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// 偷第0个房子
	x1, y1, z1 := nums[0], nums[0], 0
	// 不偷第0个房子
	x0, y0, z0:= 0, nums[1], 0
	for i := 2; i < len(nums); i++ {
		if i == len(nums) - 1 {
			z1 = y1
			z0 = max(x0 + nums[i], y0)
			break
		}
		z1 = max(x1 + nums[i], y1)
		x1, y1 = y1, z1
		z0 = max(x0 + nums[i], y0)
		x0, y0 = y0, z0
	}
	return max(z0, z1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
