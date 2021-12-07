package ac1

// 动态规划
// 设DPmax[i]表示下标i结尾的乘积最大值
// DPmin[i]表示下标i结尾的乘积最小值
// 则可以用类似lc53的方法得出状态转移方程
// DPmax[i] = max(DPmax[i-1]*nums[i], DPmin[i-1]*nums[i], nums[i])
// DPmin[i] = min(DPmax[i-1]*nums[i], DPmin[i-1]*nums[i], nums[i])
func maxProduct(nums []int) int {
	res := nums[0]
	max := nums[0]
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		max, min = getMax(getMax(max*nums[i], min*nums[i]), nums[i]),
			getMin(getMin(max*nums[i], min*nums[i]), nums[i])
		res = getMax(max, res)
	}
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
