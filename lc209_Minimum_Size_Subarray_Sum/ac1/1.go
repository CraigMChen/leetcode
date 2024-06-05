package ac1

// 二分查找
// 遍历数组 nums，对每个下标 i，利用二分查找求出最小的 j，使得 nums[i, j] 的和大于等于 target
func minSubArrayLen(target int, nums []int) int {
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	ans := 100001
	for i := 0; i < len(nums); i++ {
		l, r := i, len(nums)
		for l < r {
			m := (r-l)>>1 + l
			if preSum[m+1]-preSum[i] < target {
				l = m + 1
			} else {
				ans = min(ans, l-i+1)
				r = m
			}
		}
	}
	if ans == 100001 {
		return 0
	}
	return ans
}
