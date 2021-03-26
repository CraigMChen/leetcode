package ac2

// map
func twoSum(nums []int, target int) []int {
	n := len(nums)
	mp := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if j, ok := mp[nums[i]]; ok {
			return []int{i, j}
		}
		mp[target - nums[i]] = i
	}
	return nil
}
