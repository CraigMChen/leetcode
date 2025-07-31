package ac3

// 分块+预处理
// 将 nums 数组分为长度为 k 的若干块
// preMax[i] 表示下标 i 所属的块中的前缀最大值
// sufMax[i] 表示下标 i 所属的块中的后缀最大值
// 滑动窗口中最后一个元素的下标为 i
// 那么滑动窗口的最大值就为 max(sufMax[i-k+1], preMax[i])
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func maxSlidingWindow(nums []int, k int) []int {
	preMax := make([]int, len(nums))
	sufMax := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i%k == 0 {
			preMax[i] = nums[i]
		} else {
			preMax[i] = max(preMax[i-1], nums[i])
		}
		j := len(nums) - i - 1
		if (j+1)%k == 0 || j == len(nums)-1 {
			sufMax[j] = nums[j]
		} else {
			sufMax[j] = max(sufMax[j+1], nums[j])
		}
	}

	ans := make([]int, len(nums)-k+1)
	for i := k - 1; i < len(nums); i++ {
		ans[i-k+1] = max(sufMax[i-k+1], preMax[i])
	}
	return ans
}
