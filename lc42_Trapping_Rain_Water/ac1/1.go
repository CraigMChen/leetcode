package ac1

// 预处理
// 设 maxLeft[i] 表示 height[0:i+1] 的最大值
// maxRight[i] 表示height[i:] 的最大值
// 则位置i可以接的雨水量为 min(maxLeft[i], maxRight[i]) - height[i]
// 将每个位置的结果累加起来即为答案
// 其中 maxLeft 可以在遍历的过程中求得， maxRight 可以通过预处理后缀数组求得
func trap(height []int) int {
	maxRight := make([]int, len(height))
	tmp := 0
	for i := len(height) - 1; i >= 0; i-- {
		tmp = max(tmp, height[i])
		maxRight[i] = tmp
	}

	ans := 0
	maxLeft := 0
	for i := 0; i < len(height); i++ {
		maxLeft = max(maxLeft, height[i])
		ans += min(maxLeft, maxRight[i]) - height[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
