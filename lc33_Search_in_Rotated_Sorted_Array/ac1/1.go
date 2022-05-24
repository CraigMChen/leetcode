package ac1

// 二分
func search(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (r-l)>>1 + l
		if nums[m] == target {
			return m
		}
		if nums[m] <= nums[r-1] { // 如果右边是有序的
			if target > nums[m] && target <= nums[r-1] { // 如果target在有序的区间内
				l = m + 1
			} else {
				r = m
			}
		} else { // 如果左边是有序的
			if target >= nums[l] && m > 0 && target <= nums[m-1] {
				r = m
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
