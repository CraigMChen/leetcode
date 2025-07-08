package ac1

// 回溯
func permute(nums []int) [][]int {
	var ans [][]int
	used := make([]bool, len(nums))

	var f func([]int, int)
	f = func(cur []int, pos int) {
		if pos == len(nums) {
			ans = append(ans, append([]int(nil), cur...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				cur[pos] = nums[i]
				f(cur, pos+1)
				used[i] = false
			}
		}
	}
	f(make([]int, len(nums)), 0)
	return ans
}
