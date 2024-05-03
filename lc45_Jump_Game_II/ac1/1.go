package ac1

// 贪心
// 遍历nums，维护一个当前能到达的最远距离 maxPos，以及本次跳跃能到达的最远距离 end
// 走到下标 end 时更新步数，并更新 end = maxPos
func jump(nums []int) int {
	maxPos, end, step := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == end {
			end = maxPos
			step++
		}
	}
	return step
}
