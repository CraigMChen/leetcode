package ac1

// 二分
func findPeakElement(nums []int) int {
	l, r := 0, len(nums) - 1 // 因为要比较nums[m]和nums[m + 1]，所以这里要-1，防止下标越界
	for l < r {
		m := (r-l)>>1 + l
		if nums[m] < nums[m + 1] { // 若nums[m] < nums[m + 1]则峰值肯定在右侧
			l = m + 1
		} else { // 否则峰值肯定在左侧
			r = m
		}
	}
	return l
}
