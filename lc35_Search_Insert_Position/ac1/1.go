package ac1

// 二分
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (r - l) / 2 + l
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
