package ac4

// 摩尔投票算法
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func majorityElement(nums []int) int {
	candidate := -1
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
