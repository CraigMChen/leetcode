package ac1

// 二分
// 数组旋转后有三种情况
// 1. nums[l] < nums[m] < nums[r] 此时最小值一定在 m 的左边
// 2. nums[r] < nums[l] < nums[m] 此时最小值一定在 m 的右边
// 3. nums[m] < nums[r] < nums[l] 此时最小值一定在 m 的左边
// 只要将nums[m] 与 nums[r]相比较即可
func findMin(nums []int) int {
	l, r := 0, len(nums)-1 // 由于要比较nums[r]，所以这里要-1
	for l < r {
		m := (r-l)>>1 + l
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
