package ac1

// 二分
func search(nums []int, target int) bool {
	l, r := 0, len(nums)
	for l < r {
		m := (r-l)>>1 + l
		if nums[m] == target {
			return true
		}
		if nums[l] == nums[m] && nums[m] == nums[r-1] { // 若三个点都相等，不好判断有序性
			// 则缩小范围
			l++
			r--
		} else if nums[m] <= nums[r-1] { // 若右边有序
			if target > nums[m] && target <= nums[r-1] {
				l = m + 1
			} else {
				r = m
			}
		} else { // 若左边有序
			if target >= nums[l] && m > 0 && target <= nums[m-1] {
				r = m
			} else {
				l = m + 1
			}
		}
	}
	return false
}
