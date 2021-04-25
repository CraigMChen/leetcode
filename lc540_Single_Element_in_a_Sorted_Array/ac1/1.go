package ac1

// 二分
// 以中间两个相同的元素为中点，把原数组分为左子数组和右子数组
// 如果如果左子数组的长度为奇数，则答案一定在左边
// 否则在右边
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)
	for l < r - 1 {
		m := (r - l) >> 1 + l
		if nums[m] == nums[m + 1] {
			if (m - l) & 1 == 0 {
				l = m + 2
			} else {
				r = m
			}
		} else if nums[m] == nums[m - 1] {
			if (m - 1 - l) & 1 == 0 {
				l = m + 1
			} else {
				r = m
			}
		} else {
			return nums[m]
		}
	}
	return nums[l]
}
