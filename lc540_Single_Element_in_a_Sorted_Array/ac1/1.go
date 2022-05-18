package ac1

// 二分查找
// 从区间[l,r)中寻找答案，m为中点
// 分情况讨论
// m 为奇数时
// 若 nums[m] == nums[m+1] ，则目标元素肯定在左边，向左移动右边界，r=m
// 否则，向右移动左边界，l=m+1
// m 为偶数时
// 若 nums[m] == nums[m-1]，则目标元素肯定在左边，向左移动右边界，r=m
// 否则，向右移动左边界，l=m+1
// 上面两步可以用异或运算来简化
// 当m为奇数时 m - 1 = m ^ 1
// 当m为偶数时 m + 1 = m ^ 1
func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (r-l)>>1 + l
		if m != len(nums)-1 && nums[m] == nums[m^1] { // 防止溢出
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
