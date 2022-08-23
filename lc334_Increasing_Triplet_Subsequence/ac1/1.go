package ac1

// 预处理
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 先预处理，求出下标 0 到 i 的最小值 preMin
// 和下标 i 到 n-1 的最大值 sufMax
// 遍历 nums，若存在 preMin[i] < nums[i] < sufMax[i] 则返回 true
// 否则返回 false
func increasingTriplet(nums []int) bool {
	preMin := make([]int, len(nums))
	sufMax := make([]int, len(nums))
	preMin[0] = nums[0]
	sufMax[len(nums)-1] = nums[len(nums)-1]

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(nums); i++ {
		preMin[i] = min(preMin[i-1], nums[i])
	}
	for i := len(nums) - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], nums[i])
	}
	for i := 1; i < len(nums)-1; i++ {
		if preMin[i-1] < nums[i] && sufMax[i+1] > nums[i] {
			return true
		}
	}
	return false
}
