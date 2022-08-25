package ac1

// 前缀和
// 用 preSumCount 维护当前前缀和出现的次数
// 遍历 nums
// preSumCount[preSum-k] 即为以 nums[i] 结尾的和为 k 的子数组个数
func subarraySum(nums []int, k int) int {
	res := 0
	preSum := 0
	preSumCount := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		res += preSumCount[preSum-k]
		preSumCount[preSum]++
	}
	return res
}
