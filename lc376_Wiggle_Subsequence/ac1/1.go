package ac1

// 动态规划
// 设 up[i] 表示以下标 i 结尾的摆动上升序列（最后一段为上升序列的摆动序列）
// down[i] 表示以下标 i 结尾的摆动下降序列（最后一段为下降序列的摆动序列），则
// up[i] = up[i-1] (nums[i] <= nums[i-1])
//       = max(up[i-1], down[i-1]+1) (nums[i] < nums[i-1])
// down[i] = down[i-1] (nums[i] >= nums[i-1])
//         = max(down[i-1], up[i-1]+1) (nums[i] < nums[i-1])
func wiggleMaxLength(nums []int) int {
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			down = max(down, up+1)
		} else if nums[i] > nums[i-1] {
			up = max(up, down+1)
		}
	}
	return max(up, down)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
