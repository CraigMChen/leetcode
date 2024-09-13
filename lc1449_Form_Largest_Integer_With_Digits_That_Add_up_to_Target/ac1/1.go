package ac1

import "math"

// 动态规划
// 定义 dp[i][j] 表示使用前 i 个数位，花费刚好为 j 时能够组成的最大位数
// 则 dp[0][0] = 0
// dp[0][j] = -♾️ (j > 0)
// 状态转移方程为 dp[i][j] = dp[i-1][j] (j < cost[i])
//
//	dp[i][j] = max(dp[i-1][j], dp[i][j-cost[i]]+1) (j >= cost[i])
//
// 最终求得的 dp[9][target] 即为最大位数
// 然后从 dp[9][target] 开始倒推出最大值
func largestNumber(cost []int, target int) string {
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = math.MinInt32
	}
	for i := 0; i < len(cost); i++ {
		for j := cost[i]; j <= target; j++ {
			dp[j] = max(dp[j], dp[j-cost[i]]+1)
		}
	}
	if dp[target] < 0 {
		return "0"
	}
	var ans []byte
	for i, j := 9, target; i > 0; i-- {
		for c := cost[i-1]; j >= c && dp[j] == dp[j-c]+1; j -= c {
			ans = append(ans, '0'+byte(i))
		}
	}
	return string(ans)
}
