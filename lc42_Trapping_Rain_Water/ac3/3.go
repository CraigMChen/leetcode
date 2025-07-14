package ac3

// 双指针
// 在方法 1 中，计算当前遍历柱子能接到的雨水时的公式为
// min(preMax[i], sufMax[i]) - height[i]
// 观察这个公式，计算结果时实际上只关心 preMax[i], sufMax[i] 两者之间的较小值
// 所以只需要知道当前柱子左边和右边之间的最大值的较小的那个
// 初始化两个指针 left = 0, right = len(height) - 1
// left 指针只能向右走，right 只能向左走
// 用 leftMax 维护 left 指针经过路径上的最大值
// 用 rightMax 维护 right 指针经过路径上的最大值
// 每次只移动对应柱子高度较低的指针
// 如 height[left] <= height[right] 时
// 一定有 leftMax <= rightMax
// 那么此时 left 指针对应柱子能够接的雨水量只与 leftMax 有关，即
// ans += leftMax - height[left]
// 右边同理
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[0], height[len(height)-1]
	ans := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] <= height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}
