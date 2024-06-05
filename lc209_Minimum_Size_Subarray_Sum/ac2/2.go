package ac2

// 二分查找
// 对结果数组的长度做二分查找，判断该长度的子数组元素和是否有可能大于 target
func minSubArrayLen(target int, nums []int) int {
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	isOk := func(l int) bool {
		if l == 0 {
			return false
		}
		for i := 0; i < len(nums)-l+1; i++ {
			if preSum[i+l]-preSum[i] >= target {
				return true
			}
		}
		return false
	}
	l, r := 0, len(nums)+1
	for l < r {
		m := (r-l)>>1 + l
		if !isOk(m) {
			l = m + 1
		} else {
			r = m
		}
	}
	if l > len(nums) {
		return 0
	}
	return l
}
