package ac1

// 计数排序
// 时间复杂度 O(n) 遍历 2 次
// 空间复杂度 O(1)
func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			red++
		}
		if nums[i] == 1 {
			white++
		}
		if nums[i] == 2 {
			blue++
		}
	}
	for i := 0; i < len(nums); i++ {
		if red > 0 {
			nums[i] = 0
			red--
		} else if white > 0 {
			nums[i] = 1
			white--
		} else {
			nums[i] = 2
			blue--
		}
	}
}
