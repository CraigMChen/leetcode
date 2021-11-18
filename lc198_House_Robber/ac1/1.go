package ac1

// 动态规划
// dp[i] 表示下标范围 [1,i] 内能偷的最大价值，则
// dp[0] = nums[0]
// dp[1] = max(nums[0], nums[1])
// dp[i] = max(nums[i-2]+nums[i], nums[i-1)
// dp[len(nums)-1] 即为答案
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	if nums[0] > nums[1] {
		nums[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		if nums[i-1] > nums[i]+nums[i-2] {
			nums[i] = nums[i-1]
		} else {
			nums[i] += nums[i-2]
		}
	}
	return nums[len(nums)-1]
}
