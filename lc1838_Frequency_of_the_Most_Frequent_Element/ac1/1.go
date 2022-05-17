package ac1

import "sort"

// 滑动窗口
// 将区间[l,r]中的元素全部变为nums[r]所需的操作数
// △(l,r) = ∑(i=l,r)(nums[r]-nums[i]) = (r-l)nums[r] - ∑(i=l,r-1)nums[i]
// 若向右滑动右边界
// △(l,r+1) - △(l,r) = (r+1-l)(nums[r+1]-nums[r]) >= 0
// 若向右滑动左边界
// △(l+1,r) - △(l,r) = -(nums[r]-nums[l]) <= 0
// 则向右滑动左边界会使结果变小
// 枚举右边界，逐渐向右滑动左边界直到结果不大于k
func maxFrequency(nums []int, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Ints(nums)

	res := 1
	l, r := 0, 1
	delta := 0
	for ; r < len(nums); r++ {
		delta += (r - l) * (nums[r] - nums[r-1])
		for delta > k {
			delta -= nums[r] - nums[l]
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
