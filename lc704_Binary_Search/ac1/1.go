package ac1

// [l,r] 左闭右闭区间
// 由于整数除法为向下取整，算得的m值可能会偏小。
// 例如当l = 5, r = 6时，m = 5会偏向l
// 如果此时需要将区间右移，则l必须为m+1，否则l会永远等于5，区间不缩小，进入死循环
//
// 当l = r时，区间[l, r]还有一个数字待查找，此时不能停止循环，所以循环的条件为l <= r
// 查找区间为左闭右闭区间，r所代表的下标是在查找范围内的，所以r的初始值为len(nums)-1
// 当区间需要左移时，区间的有边界变为m的前一个值，即[l, m-1]，所以r = m - 1
func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		m := (r - l) / 2 + l
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}