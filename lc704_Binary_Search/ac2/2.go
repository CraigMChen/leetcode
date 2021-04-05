package ac2

// [l, r) 左闭右开区间
// 当l=r时，区间[l,r)中已经没有数字了，此时循环应该结束，所以循环继续的条件为l < r
// 查找区间为左闭右开区间，r所表示的下标不在查找范围内，所以r的初始值为len(nums)
// 当区间需要左移时，要将区间的右边界变为m的前一个值，[l, m-1]，即[l, m)，所以此时r = m
func search(nums []int, target int) int {
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
	return -1
}
