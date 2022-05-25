package ac1

// 二分查找
// 类似 lc153
// 当 nums[m] > nums[r] 时，最小值一定在右边
// 当 nums[m] < nuns[r] 时，最小值一定在左边
// 这题多了一种 nums[m] == nums[r] 的情况
// 此时只能将右边界左移一个单位，不能右移左边界，因为左边界不参与比较，会导致最小值直接被跳过
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (r-l)>>1 + l
		if nums[m] > nums[r] {
			l = m + 1
		} else if nums[m] < nums[r] {
			r = m
		} else {
			r--
		}
	}
	return nums[l]
}
