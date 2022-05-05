package ac1

import "sort"

// 排序+二分查找
// 先对数组进行排序
// 枚举第一条边和第二条边，二分查找大于等于第一条边和第二条边的和的元素下标，该下边之前的元素都可以作为第三条边
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			l, r := j+1, len(nums)
			for l < r {
				m := (r-l)>>1 + l
				if nums[m] < nums[i]+nums[j] {
					l = m + 1
				} else {
					r = m
				}
			}
			if l >= j+1 {
				res += l - 1 - j
			}
		}
	}
	return res
}
