package ac1

func wiggleMaxLength(nums []int) int {
	up := make([]int, len(nums))
	down := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		up[i] = 1
		down[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				up[i] = max(up[i], down[j] + 1)
			}
			if nums[i] < nums[j] {
				down[i] = max(down[i], up[j] + 1)
			}
		}
		res = max(max(up[i], down[i]), res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
