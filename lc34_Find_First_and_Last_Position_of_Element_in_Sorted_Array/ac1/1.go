package ac1

// 二分
func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)
	for l < r {
		m := (r-l)>>1 + l
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}
	left := l

	l, r = 0, len(nums)
	for l < r {
		m := (r-l)>>1 + l
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	right := l - 1 // 找右边界时要减1

	return []int{left, right}
}
