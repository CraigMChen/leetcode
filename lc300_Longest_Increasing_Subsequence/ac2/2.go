package ac2

// 贪心，二分
// 设d[i]表示长度为i的最长上升子序列的结尾最小的数
func lengthOfLIS(nums []int) int {
	var d []int
	d = append(d, -1)
	d = append(d, nums[0])
	for i := 0; i < len(nums); i++ {
		if nums[i] > d[len(d) - 1] {
			d = append(d, nums[i])
		} else {
			l, r := 1, len(d)
			for l < r {
				m := (r - l) >> 1 + l
				if d[m] < nums[i] {
					l = m + 1
				} else {
					r = m
				}
			}
			d[l] = nums[i]
		}
	}
	return len(d) - 1
}
