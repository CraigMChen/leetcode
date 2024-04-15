package ac1

// 辅助数组
func rotate(nums []int, k int) {
	tmp := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
}
