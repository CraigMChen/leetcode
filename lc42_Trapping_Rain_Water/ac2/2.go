package ac2

// 单调栈
// 维护一个单调栈，栈中的元素表示 height 的下标，且从栈低到栈顶的元素满足 height[i] 递减
// 遍历 height，若当前高度 height[i] 大于栈顶下标 top 对应的高度，则当前的柱子和栈顶元素的下一个元素 left 柱子组成的容器可以接到雨水
// 接到雨水的量为 (min(height[top], height[left]) - height[top]) * (i - left - 1_
func trap(height []int) int {
	stack := make([]int, 0, len(height))
	ans := 0
	i := 0
	for i < len(height) {
		if len(stack) == 0 || height[i] <= height[stack[len(stack)-1]] { // 栈为空或当前柱子小于等于栈顶下标对应的柱子
			stack = append(stack, i) // 当前下标入栈，维护栈的单调性
			i++
			continue
		}
		// 当前柱子大于栈顶元素对应柱子的高度
		top := stack[len(stack)-1] // 栈顶元素出栈
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			left := stack[len(stack)-1] // top 下面的元素对应的柱子，一定比 top 对应的柱子高（或相等），与当前柱子可以接到雨水
			ans += (min(height[left], height[i]) - height[top]) * (i - left - 1)
		}
	}
	return ans
}
