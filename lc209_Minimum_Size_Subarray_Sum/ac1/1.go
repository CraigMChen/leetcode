package ac1

// 二分
func minSubArrayLen(target int, nums []int) int {
	sums := make([]int, len(nums)+1)
	for i, n := range nums {
		sums[i+1] = sums[i] + n
	}
	if sums[len(nums)] < target {
		return 0
	}
	ans := len(nums)
	for i := 0; i < len(nums); i++ { // 遍历所有下标i
		l, r := i, len(nums)
		for l < r { // 对于每个下标i，二分搜索求出最小的x，使得nums[x]的前缀和大于等于target
			m := (r-l)>>1 + l
			if sums[m+1]-sums[i] >= target {
				if m-i+1 < ans {
					ans = m - i + 1
				}
				r = m
			} else {
				l = m + 1
			}
		}
	}
	return ans
}
