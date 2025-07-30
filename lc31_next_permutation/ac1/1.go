package ac1

// 更新策略：找到一个较小值和较大值，较小值在较大值的左边
// 且较小值尽可能靠右，较大值尽可能小，交换这两个值，并把原较小值下标后面的元素倒序排列
// 从右往左，找到第一个 i，使得 nums[i] < nums[i+1]，nums[i] 就是较小值
// 若找不到，说明原数组是按照降序排列的，直接返回数组的倒序
// 否则，用二分查找，找出 nums[i+1:] 中第一个小于等于 nums[i] 的数 nums[l]
// 那么 nums[l-1] 就是较大值
// 交换 nums[i] 和 nums[l-1] 并倒序 nums[i+1:] 即可
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i >= 0 {
		l, r := i+1, len(nums)
		for l < r {
			m := (r-l)>>1 + l
			if nums[m] > nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}
		nums[i], nums[l-1] = nums[l-1], nums[i]
	}
	for j, k := i+1, len(nums)-1; j < k; j, k = j+1, k-1 {
		nums[j], nums[k] = nums[k], nums[j]
	}
}
