package ac1

// 回溯
// 时间复杂度 O(n*2^n)
// 空间复杂度 O(n)
func subsets(nums []int) [][]int {
	var ans [][]int
	var backtrack func([]int, int)
	backtrack = func(cur []int, index int) {
		ans = append(ans, append([]int{}, cur...))
		for i := index; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack(cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack([]int{}, 0)
	return ans
}
