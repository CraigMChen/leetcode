package ac1

// 二分
// 设count[i]表示nums中小于等于i的数的个数，target为目标值
// 则对于所有[1, target)都满足count[i] <= i
// 对于所有[target, n - 1]都满足count[i] > i
// 则可以进行二分搜索，复杂度为O(nlogn)
func findDuplicate(nums []int) int {
	getCount := func(m int) int {
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= m {
				count++
			}
		}
		return count
	}
	l, r := 1, len(nums)
	for l < r {
		m := (r-l)>>1 + l
		if getCount(m) > m {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
