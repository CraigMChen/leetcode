package ac1

// 预处理
// 分别求出前缀最大值 preMax 和 后缀最大值 sufMax
// 则位置 i 能够接到的雨水为 max(min(preMax[i-1], sufMax[i+1])-height[i], 0)
func trap(height []int) int {
	sufMax := make([]int, len(height)+1)
	for i := len(height) - 1; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}
	preMax, ans := 0, 0
	for i := 0; i < len(height); i++ {
		ans += max(min(preMax, sufMax[i+1])-height[i], 0)
		preMax = max(preMax, height[i])
	}
	return ans
}
