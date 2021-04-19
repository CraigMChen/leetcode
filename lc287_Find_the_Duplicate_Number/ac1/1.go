package ac1

// 二分
// 设count[i]表示nums中小于等于i的数的个数，target为目标值
// 则对于所有[1, target)都满足count[i] <= i
// 对于所有[target, n - 1]都满足count[i] > i
// 则可以进行二分搜索，复杂度为O(nlogn)

func findDuplicate(nums []int) int {
	l, r := 0, len(nums)
	ans := -1
	for l < r {
		m := (r - l) >> 1 + l
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= m {
				count++
			}
		}
		if count <= m {
			l = m + 1
		} else {
			ans = m
			r = m
		}
	}
	return ans
}
