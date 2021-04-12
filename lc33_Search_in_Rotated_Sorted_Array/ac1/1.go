package ac1

// 二分
func search(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := ((r - l) >> 1) + l
		if nums[m] == target {
			return m
		} else if nums[m] <= nums[r - 1] { // 如果右边是有序的
			if target > nums[m] && target <= nums[r - 1] { // 且target在(nums[m],nums[r-1]]范围内
				l = m + 1 // 则target肯定在右边
			} else {
				r = m // 否则target一定在左边
			}
		} else { // 如果左边是有序的
			if target >= nums[l] && target < nums[m] { // 且target在[nums[l],nums[m])范围内
				r = m // 则target肯定在左边
			} else {
				l = m + 1 // 否则target一定在右边
			}
		}
	}
	if l == len(nums) || nums[l] != target {
		return -1
	}
	return l
}
