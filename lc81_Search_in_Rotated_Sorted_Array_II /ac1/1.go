package ac1

// 二分
func search(nums []int, target int) bool {
	l, r := 0, len(nums)
	for l < r {
		m := (r-l)>>1 + l
		if nums[m] == target || nums[l] == target || nums[r-1] == target {
			return true
		} else if nums[m] == nums[l] {
			l++
		} else if nums[m] == nums[r-1] {
			r--
		} else if nums[m] <= nums[r-1] {
			if target > nums[m] && target <= nums[r-1] {
				l = m + 1
			} else {
				r = m
			}
		} else {
			if target >= nums[l] && target < nums[m] {
				r = m
			} else {
				l = m + 1
			}
		}
	}
	if l == len(nums) || nums[l] != target {
		return false
	}
	return true
}
