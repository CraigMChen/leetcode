package ac1

// 暴力
func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n - 1; i++ {
		x := target - nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] == x {
				return []int{i, j}
			}
		}
	}
	return nil
}
