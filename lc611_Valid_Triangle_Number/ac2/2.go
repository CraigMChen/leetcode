package ac2

import "sort"

// 双指针
// 首先将数组排序，三条边的下标 i, j, k 满足
// nums[i] + nums[j] > nums[k] (i < j < k)
// 对于一个固定的 i，当 j 增加时，k 一定要增加才能满足上面的不等式
// 枚举 i，用双指针法求 j 和 k
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		k := i + 1
		for j := i + 1; j < len(nums); j++ {
			for k < len(nums) && nums[k] < nums[i]+nums[j] {
				k++
			}
			res += max(k-1-j, 0)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
