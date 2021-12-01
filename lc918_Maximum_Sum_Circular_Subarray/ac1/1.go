package ac1

// 动态规划+邻接数组
// 循环数组的子段可以是单区间子段或双区间子段
// 对于单区间子段的最大和可以像lc53一样求得
// 对于双区间子段[0,i],[j,len(nums)-1]，可以用以下方法求得：
// 设leftSum表示nums数组的前缀和
// 设rightSum表示数组nums的后缀和
// 设maxRightSum[j]=max(rightSum[j,len(nums)-1])
// 则双区间子段问题可以转化为，求最大的i，使得
// leftSum[i] + maxRightSum[i+1] 最大
// 表达式中的两项都可以通过预处理求得
// 则可以遍历i，求出表达式的最大值，即为双区间子段的最大和
// 最终结果为单区间子段的最大和与双区间子段的最大和二者的较大者
func maxSubarraySumCircular(nums []int) int {
	// 先求单区间子段最大和
	res, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur = getMax(cur, 0) + nums[i]
		res = getMax(res, cur)
	}

	// 再求双区间子段最大和
	leftSum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		leftSum = append(leftSum, leftSum[i-1]+nums[i])
	}

	rightSum := 0
	maxRightSum := make([]int, len(nums)+1)
	max := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		rightSum += nums[i]
		max = getMax(max, rightSum)
		maxRightSum[i] = max
	}

	for i := 0; i < len(nums); i++ {
		cur = leftSum[i] + maxRightSum[i+1]
		res = getMax(res, cur)
	}
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}