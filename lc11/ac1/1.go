package ac1

// 双指针
func maxArea(height []int) int {
	ans := 0
	l := len(height)
	for i, j := 0, l - 1; i < j; {
		s := 0
		if height[i] < height[j] {
			s = height[i] * (j - i)
			i++
		} else {
			s = height[j] * (j - i)
			j--
		}
		if s > ans {
			ans = s
		}
	}
	return ans
}
