package ac2

// 数组置换
// 将 nums 数组中的元素置换位置，使得置换后的 nums 数组尽量满足下标 i 的元素为 i+1，即 nums[i] = i+1
// 那么第一个不满足 nums[i] = i+1 的 i+1 即为没有在数组中出现的第一个正整数
// 遍历 nums，若 nums[i] 在 [1,n] 范围内，则将 nums[i] 换到下标 nums[i]-1 的位置处，并把 nums[i]-1 换回来
// 当 nums[i] == nums[nums[i]-1] 时，表明 nums[i] 这个数已经在正确的位置上了，此时可以停止交换
// 所有元素交换完成后遍历数组，第一个 nums[i] != i+1 的 i+1 值就是没有在数组中出现的最小正整数
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
