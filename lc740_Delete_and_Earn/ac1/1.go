package ac1

//func deleteAndEarn(nums []int) int {
//	sum := make(map[int]int)
//	max := 0
//	min := 20001
//	for i := range nums {
//		if nums[i] > max {
//			max = nums[i]
//		}
//		if nums[i] < min {
//			min = nums[i]
//		}
//		sum[nums[i]]++
//	}
//	x, y, z := 0, 0, 0
//	for i := min; i <= max; i++ {
//		z = getMax(x + i * sum[i], y)
//		x, y = y, z
//	}
//	return z
//}

// 动态规划
// 设sum[i]表示i在nums中的出现次数
// dp[i]表示删除i获得的最大收益，则
// dp[i] = max(dp[i-2]+sum[i]*i, dp[i-1])
// dp[max(nums)]即为最终结果
func deleteAndEarn(nums []int) int {
	max := 0
	min := 20001
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	sum := make([]int, max - min + 1)
	for i := range nums {
		sum[nums[i]-min]++
	}
	x, y, z := 0, 0, 0
	for i := min; i <= max; i++ {
		z = getMax(x + i * sum[i - min], y)
		x, y = y, z
	}
	return z
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
