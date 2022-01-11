package ac1

// 动态规划
// 设 dp[i] 表示以 nums[i-1], nums[i] 两个元素结尾的等差数列的个数
// 则 当 nums[i] - nums[i-1] == nums[i-1] - nums[i-2] 时，dp[i] = dp[i-1] + 1
// 否则 dp[i] = 0
// 答案为所有 dp 之和
// 初始状态：
// 若 nums[0:3] 为等差数列，则 dp[2] = 1，否则 dp[2] = 0
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	res := 0
	last := 0
	d := nums[2] - nums[1]
	if nums[1] - nums[0] == d {
		last = 1
		res = 1
	}
	for i := 3; i < len(nums); i++ {
		if nums[i] - nums[i-1] == d {
			last++
			res += last
		} else {
			last = 0
		}
		d = nums[i] - nums[i-1]
	}
	return res
}
