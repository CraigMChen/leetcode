package ac2

// 贪心
// maxIndex表示当前最大能到达的位置
// 遍历nums，更新当前最大可达位置
// 遍历过程中若当前位置比最大可达位置大，则不能走完
// 遍历过程中若最大可达位置大于等于nums长度，则能走完
func canJump(nums []int) bool {
	maxIndex := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxIndex < i {
			return false
		}
		maxIndex = max(maxIndex, i+nums[i])
	}
	return true
}
