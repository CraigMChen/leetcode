package ac2

// 二分
// lis[i] 表示长度为 i+1 的最长递增子序列的末尾元素的最小值
// 易知，lis 一定是递增的
// 对每个 nums[i] ，用二分查找找出 lis 中第一个大于 nums[i] 的元素下标，并更新该位置为 nums[i]
func lengthOfLIS(nums []int) int {
	lis := make([]int, len(nums))
	lis[0] = nums[0]
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > lis[length-1] {
			lis[length] = nums[i]
			length++
			continue
		}
		l, r := 0, length
		for l < r {
			m := (r-l)>>1 + l
			if lis[m] < nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}
		lis[l] = nums[i]
	}
	return length
}
