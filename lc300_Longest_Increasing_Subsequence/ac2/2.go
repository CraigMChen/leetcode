package ac2

// 二分
// 用数组 lis 表示 nums 的最长递增子序列
// 对每个 nums[i] ，用二分查找找出能够在 lis 中插入的下标并替换
func lengthOfLIS(nums []int) int {
	var lis []int
	for i := 0; i < len(nums); i++ {
		l, r := 0, len(lis)
		for l < r {
			m := (r-l)>>1 + l
			if lis[m] < nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}
		if l == len(lis) {
			lis = append(lis, nums[i])
		} else {
			lis[l] = nums[i]
		}
	}
	return len(lis)
}
