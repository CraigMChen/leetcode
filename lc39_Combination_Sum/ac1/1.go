package ac1

// 回溯
// 时间复杂度 O(S)，其中 S 为所有可行解的长度之和
// 空间复杂度 O(target)，空间复杂度取决于答案数组的长度和递归层数，最差情况下需要递归 target 层
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var com func(cur []int, sum int, i int)
	com = func(cur []int, sum int, i int) {
		if sum == target {
			ans = append(ans, append([]int{}, cur...))
			return
		}
		if sum > target {
			return
		}
		for j := i; j < len(candidates); j++ {
			cur = append(cur, candidates[j])
			com(cur, sum+candidates[j], j)
			cur = cur[:len(cur)-1]
		}
	}
	com([]int{}, 0, 0)
	return ans
}
