package ac3

// 双指针
// 时间复杂度 O(n) 遍历 1 次
// 空间复杂度 O(1)
// 遍历 nums
// 若 nums[i] 为 0，将 nums[i] 与 nums[p0] 交换位置。若 p0 < p1，此时一定将一个 1 换到了 nums[1]，所以要换回来
// 		分别将 p0 和 p1 向右移动一个单位
// 若 nums[i] 为 1，将 nums[i] 与 nums[p1] 交换位置，并将 p1 向右移动一个单位
func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
