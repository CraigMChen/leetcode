package ac4

// 双指针
// 时间复杂度 O(n) 遍历一次
// 空间复杂度 O(1)
// 遍历 nums，当 nums[i] 为 0 时把 nums[i] 与 nums[p0] 交换位置，并把 p0 向右移动一个单位
// 当 nums[i] 为 2 时把 nums[i] 与 nums[p2] 交换位置，并把 p2 向左移动一个单位
// 此时若 nums[i] 还是 2，就不断重复上一步
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
