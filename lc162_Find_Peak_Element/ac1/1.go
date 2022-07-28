package ac1

// 二分
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (r-l)>>1 + l
		if m < len(nums)-1 && nums[m] < nums[m+1] { // 如果 m 不是最后有一个元素且 nuns[m] < nums[m+1]
			l = m + 1 // 则右侧一定存在峰值
		} else { // 否则
			r = m // 左侧一定存在峰值
		}
	}
	return l
}
