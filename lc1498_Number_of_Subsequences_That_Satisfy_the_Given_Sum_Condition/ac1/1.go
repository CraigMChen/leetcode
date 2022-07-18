package ac1

import "sort"

const MOD int = 1e9 + 7

// 二分查找
// 首先将数组排序
// 从头开始遍历数组，对于每个元素i，使用二分查找找到第一个下标j使得nums[i]+nums[j]>target
// 则区间[i,j)中，符合要求的子数组个数为 2^(j-1-1)，因为
// 元素nums[i] 必选，后面[i+1,j)的每个元素都有存在和不存在两种状态
// 把所有结果累加即为答案
func numSubseq(nums []int, target int) int {
	pow := func(x, n int) int {
		res, y := 1, x
		for n > 0 {
			if n&1 == 1 {
				res = res * y % MOD
			}
			y = y * y % MOD
			n >>= 1
		}
		return res
	}
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		l, r := i, len(nums)
		for l < r {
			m := (r-l)>>1 + l
			if nums[m]+nums[i] > target {
				r = m
			} else {
				l = m + 1
			}
		}
		if l != i {
			res = (res + pow(2, l-i-1)) % MOD // nums[i] 必选，所以是 2^(l-i-1)
		}
	}
	return res % MOD
}
