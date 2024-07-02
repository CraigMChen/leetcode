package ac1

// 滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = struct{}{}
	}
	return false
}
