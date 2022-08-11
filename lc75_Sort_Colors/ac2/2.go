package ac2

// 单指针
// 时间复杂度 O(n) 遍历 2 次
// 空间复杂度 O(1)
func sortColors(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}
	for i := k; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}
}
