package ac3

// 滑动窗口
// 窗口从 [0, 0] 开始
// 若窗口中元素和偏小，则右边界往右移动
// 若窗口中元素和偏大，则左边界往右移动
// 若窗口只有 1 个元素，则左边界和右边界都往右移动
func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := nums[0]
	ans := 100001
	for l < len(nums) && r < len(nums) {
		if sum >= target {
			ans = min(ans, r-l+1)
			if l == r {
				r++
				if r < len(nums) {
					sum += nums[r]
				}
			}
			sum -= nums[l]
			l++
		} else {
			r++
			if r < len(nums) {
				sum += nums[r]
			}
		}
	}
	if ans == 100001 {
		return 0
	}
	return ans
}
